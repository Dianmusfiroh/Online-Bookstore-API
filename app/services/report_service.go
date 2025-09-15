package services

import (
    "book-online-api/app/models"
    "book-online-api/app/repository"
)

type ReportService interface {
    GetSalesReport() (float64, int64, error)
    GetBestsellers() ([]models.BestsellerBook, error)
    GetPriceReport() (*models.PriceReport, error)
}

type reportService struct {
    repo repository.ReportRepository
}

func NewReportService(repo repository.ReportRepository) ReportService {
    return &reportService{repo}
}

func (s *reportService) GetSalesReport() (float64, int64, error) {
    return s.repo.GetSalesReport()
}

func (s *reportService) GetBestsellers() ([]models.BestsellerBook, error) {
    return s.repo.GetBestsellers()
}

func (s *reportService) GetPriceReport() (*models.PriceReport, error) {
    return s.repo.GetPriceReport()
}