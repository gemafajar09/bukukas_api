package handler

import (
	"go-project/internal/domain"
	"go-project/internal/usecase"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type CashHandler struct {
	uc usecase.CashUsecase
}

func NewCashHandler(uc usecase.CashUsecase) *CashHandler {
	return &CashHandler{uc: uc}
}

// CreateTransaction godoc
// @Summary Tambah transaksi kas
// @Description Tambah transaksi uang masuk atau keluar (requires JWT token)
// @Tags Cash
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param transaction body domain.CashTransaction true "Data transaksi kas"
// @Success 201 {object} map[string]interface{} "Transaction recorded successfully"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Router /cash-transactions [post]
func (h *CashHandler) CreateTransaction(c *gin.Context) {
	var transaction domain.CashTransaction
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.uc.RecordTransaction(transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Transaction recorded successfully"})
}

// GetTransactions godoc
// @Summary Ambil daftar transaksi kas
// @Description Ambil daftar transaksi kas berdasarkan rentang tanggal
// @Tags Cash
// @Security BearerAuth
// @Produce json
// @Param start query string false "Tanggal mulai (YYYY-MM-DD)"
// @Param end query string false "Tanggal akhir (YYYY-MM-DD)"
// @Success 200 {object} map[string]interface{} "List transaksi"
// @Failure 500 {object} map[string]interface{} "Server Error"
// @Router /cash-transactions [get]
func (h *CashHandler) GetTransactions(c *gin.Context) {
	startStr := c.Query("start")
	endStr := c.Query("end")

	start, _ := time.Parse("2006-01-02", startStr)
	end, _ := time.Parse("2006-01-02", endStr)
	if end.IsZero() {
		end = time.Now()
	}

	data, err := h.uc.GetReport(start, end)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"transactions": data})
}

// GetBalance godoc
// @Summary Lihat saldo kas harian
// @Description Menghitung dan menampilkan saldo kas untuk tanggal tertentu
// @Tags Cash
// @Security BearerAuth
// @Produce json
// @Param date query string false "Tanggal (YYYY-MM-DD)"
// @Success 200 {object} map[string]interface{} "Data saldo kas harian"
// @Failure 500 {object} map[string]interface{} "Server Error"
// @Router /cash-balance [get]
func (h *CashHandler) GetBalance(c *gin.Context) {
	dateStr := c.Query("date")
	date, _ := time.Parse("2006-01-02", dateStr)
	if date.IsZero() {
		date = time.Now()
	}

	balance, err := h.uc.CalculateDailyBalance(date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"balance": balance})
}

// GetCategories godoc
// @Summary Ambil daftar kategori kas
// @Description Menampilkan semua kategori transaksi kas (uang masuk / keluar)
// @Tags Cash
// @Security BearerAuth
// @Produce json
// @Success 200 {object} map[string]interface{} "List kategori kas"
// @Failure 500 {object} map[string]interface{} "Server Error"
// @Router /cash-categories [get]
func (h *CashHandler) GetCategories(c *gin.Context) {
	cats, err := h.uc.GetCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"categories": cats})
}
