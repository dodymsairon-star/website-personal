package handlers

import (
	"database/sql"
	"fintrack/config"
	"fintrack/models"
	"fmt"
	"net/http"
	"strconv" // Sekarang package ini AKAN DIPAKAI, jadi tidak akan error merah
	"time"

	"github.com/gin-gonic/gin"
)

// GetAllTransactions - GET /api/transactions
func GetAllTransactions(c *gin.Context) {
	rows, err := config.DB.Query(`
		SELECT id, ref, date, description, category, amount, type, status, created_at, updated_at 
		FROM transactions 
		ORDER BY date DESC
	`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var transactions []models.Transaction
	for rows.Next() {
		var tx models.Transaction
		err := rows.Scan(&tx.ID, &tx.Ref, &tx.Date, &tx.Description, &tx.Category,
			&tx.Amount, &tx.Type, &tx.Status, &tx.CreatedAt, &tx.UpdatedAt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		transactions = append(transactions, tx)
	}

	c.JSON(http.StatusOK, gin.H{"data": transactions})
}

// GetTransaction - GET /api/transactions/:id
func GetTransaction(c *gin.Context) {
	idStr := c.Param("id")

	// ✅ PERBAIKAN: Gunakan strconv untuk mengubah string menjadi angka (uint64)
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format ID tidak valid"})
		return
	}

	var tx models.Transaction
	err = config.DB.QueryRow(`
		SELECT id, ref, date, description, category, amount, type, status, created_at, updated_at 
		FROM transactions WHERE id = ?
	`, id).Scan(&tx.ID, &tx.Ref, &tx.Date, &tx.Description, &tx.Category,
		&tx.Amount, &tx.Type, &tx.Status, &tx.CreatedAt, &tx.UpdatedAt)

	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": tx})
}

// CreateTransaction - POST /api/transactions
func CreateTransaction(c *gin.Context) {
	var input models.CreateTransactionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate reference number
	var count int
	err := config.DB.QueryRow("SELECT COUNT(*) FROM transactions").Scan(&count)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ref := fmt.Sprintf("TRX%03d", count+1)

	// Set default status
	if input.Status == "" {
		input.Status = "success"
	}

	// Insert to database
	result, err := config.DB.Exec(`
		INSERT INTO transactions (ref, date, description, category, amount, type, status) 
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`, ref, input.Date, input.Description, input.Category, input.Amount, input.Type, input.Status)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	insertedID, _ := result.LastInsertId()

	// Get created transaction
	var tx models.Transaction
	config.DB.QueryRow(`
		SELECT id, ref, date, description, category, amount, type, status, created_at, updated_at 
		FROM transactions WHERE id = ?
	`, insertedID).Scan(&tx.ID, &tx.Ref, &tx.Date, &tx.Description, &tx.Category,
		&tx.Amount, &tx.Type, &tx.Status, &tx.CreatedAt, &tx.UpdatedAt)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Transaction created successfully",
		"data":    tx,
	})
}

// UpdateTransaction - PUT /api/transactions/:id
func UpdateTransaction(c *gin.Context) {
	idStr := c.Param("id")

	// ✅ PERBAIKAN: Gunakan strconv di sini juga
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format ID tidak valid"})
		return
	}

	var input models.UpdateTransactionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if exists
	var exists int
	err = config.DB.QueryRow("SELECT COUNT(*) FROM transactions WHERE id = ?", id).Scan(&exists)
	if err != nil || exists == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}

	// Update
	_, err = config.DB.Exec(`
		UPDATE transactions 
		SET date = COALESCE(?, date), 
			description = COALESCE(?, description), 
			category = COALESCE(?, category), 
			amount = COALESCE(?, amount), 
			type = COALESCE(?, type), 
			status = COALESCE(?, status)
		WHERE id = ?
	`, input.Date, input.Description, input.Category, input.Amount, input.Type, input.Status, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Get updated transaction
	var tx models.Transaction
	config.DB.QueryRow(`
		SELECT id, ref, date, description, category, amount, type, status, created_at, updated_at 
		FROM transactions WHERE id = ?
	`, id).Scan(&tx.ID, &tx.Ref, &tx.Date, &tx.Description, &tx.Category,
		&tx.Amount, &tx.Type, &tx.Status, &tx.CreatedAt, &tx.UpdatedAt)

	c.JSON(http.StatusOK, gin.H{
		"message": "Transaction updated successfully",
		"data":    tx,
	})
}

// DeleteTransaction - DELETE /api/transactions/:id
func DeleteTransaction(c *gin.Context) {
	idStr := c.Param("id")

	// ✅ PERBAIKAN: Gunakan strconv di sini juga
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format ID tidak valid"})
		return
	}

	_, err = config.DB.Exec("DELETE FROM transactions WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Transaction deleted successfully"})
}

// GetStats - GET /api/stats
func GetStats(c *gin.Context) {
	var totalIncome, totalExpense float64

	err := config.DB.QueryRow(`
		SELECT COALESCE(SUM(amount), 0) FROM transactions WHERE type = 'income'
	`).Scan(&totalIncome)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = config.DB.QueryRow(`
		SELECT COALESCE(SUM(amount), 0) FROM transactions WHERE type = 'expense'
	`).Scan(&totalExpense)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	balance := totalIncome - totalExpense
	savings := balance * 0.5

	c.JSON(http.StatusOK, gin.H{
		"data": models.Stats{
			TotalIncome:  totalIncome,
			TotalExpense: totalExpense,
			Balance:      balance,
			Savings:      savings,
		},
	})
}

// GetCategoryStats - GET /api/category-stats
func GetCategoryStats(c *gin.Context) {
	txType := c.Query("type")

	var rows *sql.Rows
	var err error

	if txType != "" {
		rows, err = config.DB.Query(`
			SELECT category, SUM(amount) as total, COUNT(*) as count 
			FROM transactions 
			WHERE type = ?
			GROUP BY category
			ORDER BY total DESC
		`, txType)
	} else {
		rows, err = config.DB.Query(`
			SELECT category, SUM(amount) as total, COUNT(*) as count 
			FROM transactions 
			GROUP BY category
			ORDER BY total DESC
		`)
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var results []models.CategoryStats
	for rows.Next() {
		var stat models.CategoryStats
		rows.Scan(&stat.Category, &stat.Total, &stat.Count)
		results = append(results, stat)
	}

	c.JSON(http.StatusOK, gin.H{"data": results})
}

// ExportCSV - GET /api/export/csv
func ExportCSV(c *gin.Context) {
	rows, err := config.DB.Query(`
		SELECT ref, date, description, category, amount, type, status 
		FROM transactions 
		ORDER BY date DESC
	`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	csv := "Referensi,Tanggal,Deskripsi,Kategori,Jumlah,Tipe,Status\n"
	for rows.Next() {
		var ref, date, desc, category, txType, status string
		var amount float64
		rows.Scan(&ref, &date, &desc, &category, &amount, &txType, &status)
		csv += fmt.Sprintf("%s,%s,%s,%s,%.0f,%s,%s\n", ref, date, desc, category, amount, txType, status)
	}

	c.Header("Content-Disposition", "attachment; filename=transactions.csv")
	c.Data(http.StatusOK, "text/csv", []byte(csv))
}

// GetMonthlyData - GET /api/monthly
func GetMonthlyData(c *gin.Context) {
	year := c.Query("year")
	if year == "" {
		year = fmt.Sprintf("%d", time.Now().Year())
	}

	months := []string{"01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12"}
	var incomeData, expenseData []float64

	for _, month := range months {
		var income, expense float64
		pattern := year + "-" + month + "%"

		config.DB.QueryRow(`
			SELECT COALESCE(SUM(amount), 0) FROM transactions 
			WHERE type = 'income' AND date LIKE ?
		`, pattern).Scan(&income)

		config.DB.QueryRow(`
			SELECT COALESCE(SUM(amount), 0) FROM transactions 
			WHERE type = 'expense' AND date LIKE ?
		`, pattern).Scan(&expense)

		incomeData = append(incomeData, income)
		expenseData = append(expenseData, expense)
	}

	c.JSON(http.StatusOK, gin.H{
		"months":  []string{"Jan", "Feb", "Mar", "Apr", "Mei", "Jun", "Jul", "Agu", "Sep", "Okt", "Nov", "Des"},
		"income":  incomeData,
		"expense": expenseData,
	})
}

// SearchTransactions - GET /api/transactions/search
func SearchTransactions(c *gin.Context) {
	query := c.Query("q")

	rows, err := config.DB.Query(`
		SELECT id, ref, date, description, category, amount, type, status, created_at, updated_at 
		FROM transactions 
		WHERE ref LIKE ? OR description LIKE ? OR category LIKE ?
		ORDER BY date DESC
	`, "%"+query+"%", "%"+query+"%", "%"+query+"%")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var transactions []models.Transaction
	for rows.Next() {
		var tx models.Transaction
		rows.Scan(&tx.ID, &tx.Ref, &tx.Date, &tx.Description, &tx.Category,
			&tx.Amount, &tx.Type, &tx.Status, &tx.CreatedAt, &tx.UpdatedAt)
		transactions = append(transactions, tx)
	}

	c.JSON(http.StatusOK, gin.H{"data": transactions})
}

// FilterTransactions - GET /api/transactions/filter
func FilterTransactions(c *gin.Context) {
	txType := c.Query("type")
	status := c.Query("status")
	category := c.Query("category")
	dateFrom := c.Query("date_from")
	dateTo := c.Query("date_to")

	query := `
		SELECT id, ref, date, description, category, amount, type, status, created_at, updated_at 
		FROM transactions WHERE 1=1
	`
	args := []interface{}{}

	if txType != "" {
		query += " AND type = ?"
		args = append(args, txType)
	}
	if status != "" {
		query += " AND status = ?"
		args = append(args, status)
	}
	if category != "" {
		query += " AND category = ?"
		args = append(args, category)
	}
	if dateFrom != "" {
		query += " AND date >= ?"
		args = append(args, dateFrom)
	}
	if dateTo != "" {
		query += " AND date <= ?"
		args = append(args, dateTo)
	}

	query += " ORDER BY date DESC"

	rows, err := config.DB.Query(query, args...)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var transactions []models.Transaction
	for rows.Next() {
		var tx models.Transaction
		rows.Scan(&tx.ID, &tx.Ref, &tx.Date, &tx.Description, &tx.Category,
			&tx.Amount, &tx.Type, &tx.Status, &tx.CreatedAt, &tx.UpdatedAt)
		transactions = append(transactions, tx)
	}

	c.JSON(http.StatusOK, gin.H{"data": transactions})
}

// GetDashboardData - GET /api/dashboard
func GetDashboardData(c *gin.Context) {
	var totalIncome, totalExpense float64
	var incomeCount, expenseCount int

	config.DB.QueryRow(`SELECT COALESCE(SUM(amount), 0) FROM transactions WHERE type = 'income'`).Scan(&totalIncome)
	config.DB.QueryRow(`SELECT COALESCE(SUM(amount), 0) FROM transactions WHERE type = 'expense'`).Scan(&totalExpense)
	config.DB.QueryRow(`SELECT COUNT(*) FROM transactions WHERE type = 'income'`).Scan(&incomeCount)
	config.DB.QueryRow(`SELECT COUNT(*) FROM transactions WHERE type = 'expense'`).Scan(&expenseCount)

	// Recent transactions
	rows, _ := config.DB.Query(`
		SELECT id, ref, date, description, category, amount, type, status, created_at, updated_at 
		FROM transactions ORDER BY date DESC LIMIT 10
	`)
	defer rows.Close()

	var recentTransactions []models.Transaction
	for rows.Next() {
		var tx models.Transaction
		rows.Scan(&tx.ID, &tx.Ref, &tx.Date, &tx.Description, &tx.Category,
			&tx.Amount, &tx.Type, &tx.Status, &tx.CreatedAt, &tx.UpdatedAt)
		recentTransactions = append(recentTransactions, tx)
	}

	// Category expenses
	catRows, _ := config.DB.Query(`
		SELECT category, SUM(amount) as total, COUNT(*) as count 
		FROM transactions WHERE type = 'expense'
		GROUP BY category ORDER BY total DESC LIMIT 5
	`)
	defer catRows.Close()

	var categoryExpenses []models.CategoryStats
	for catRows.Next() {
		var stat models.CategoryStats
		catRows.Scan(&stat.Category, &stat.Total, &stat.Count)
		categoryExpenses = append(categoryExpenses, stat)
	}

	c.JSON(http.StatusOK, gin.H{
		"stats": models.Stats{
			TotalIncome:  totalIncome,
			TotalExpense: totalExpense,
			Balance:      totalIncome - totalExpense,
			Savings:      (totalIncome - totalExpense) * 0.5,
		},
		"counts": gin.H{
			"income":  incomeCount,
			"expense": expenseCount,
		},
		"recent":           recentTransactions,
		"category_expense": categoryExpenses,
	})
}
