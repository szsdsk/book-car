package database

import (
	"acs/src/models"
	"github.com/tealeg/xlsx"
	"log"
	"time"
)

func initData() {
	xl, err := xlsx.OpenFile("./data/CarShareReservation.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	headerMap := make(map[string]int)
	for i, cell := range xl.Sheets[0].Rows[0].Cells {
		headerMap[cell.String()] = i
	}
	var records []models.BookRecord
	for _, sheet := range xl.Sheets {
		for rowIndex, row := range sheet.Rows {
			if rowIndex == 0 {
				continue
			}
			cells := row.Cells
			carID, _ := cells[headerMap["Car ID"]].Int()
			capacity, _ := cells[headerMap["Number of Passengers"]].Int()
			pricePerHour, _ := cells[headerMap["Price Per Hour"]].Float()
			pricePerDay, _ := cells[headerMap["Price Per Day"]].Float()
			car := models.Car{
				Id:           carID,
				Make:         cells[headerMap["Make"]].String(),
				Model:        cells[headerMap["Model"]].String(),
				PricePerHour: pricePerHour,
				PricePerDay:  pricePerDay,
				Capacity:     capacity,
				Description:  cells[headerMap["Description"]].String(),
				Img:          cells[headerMap["Img"]].String(),
			}
			locationId, _ := cells[headerMap["Location ID"]].Int()
			location := models.Location{
				Id:            locationId,
				StreetAddress: cells[headerMap["Street Address"]].String(),
				Telephone:     cells[headerMap["Locations_Telephone"]].String(),
			}
			tickets, _ := cells[headerMap["Tickets (last 3 years)"]].Int()
			license, _ := cells[headerMap["License Number"]].Int()
			student := cells[headerMap["Student?"]].Bool()
			customer := models.Customer{
				Uid:        cells[headerMap["Customer ID"]].String(),
				Firstname:  cells[headerMap["First Name"]].String(),
				Lastname:   cells[headerMap["Last Name"]].String(),
				Address:    cells[headerMap["Address"]].String(),
				Email:      cells[headerMap["Email Address"]].String(),
				CreditCard: cells[headerMap["Credit Card"]].String(),
				IsStudent:  &student,
				Telephone:  cells[headerMap["Customers_Telephone"]].String(),
				Phone:      cells[headerMap["Cell Phone"]].String(),
				Licence:    license,
				Tickets:    tickets,
				StateIssue: cells[headerMap["State Issued"]].String(),
			}
			var expiration time.Time
			if cells[headerMap["Expiration Date"]].String() != "" {
				expiration, _ = time.ParseInLocation("2006/01/02", cells[headerMap["Expiration Date"]].String(), time.Local)
			}
			customer.Expiration = expiration
			pickUpTime, _ := time.ParseInLocation("2\\-Jan\\-2006\\ 15:04:05", cells[headerMap["Pickup"]].String(), time.Local)
			dropOff, _ := time.ParseInLocation("2\\-Jan\\-2006\\ 15:04:05", cells[headerMap["Drop Off"]].String(), time.Local)
			reservedDate, _ := time.ParseInLocation("2\\-Jan\\-2006", cells[headerMap["Date Reserved"]].String(), time.Local)
			bookRecord := models.BookRecord{
				ReservationNum: cells[headerMap["Reservation Number"]].String(),
				PricePerHour:   pricePerHour,
				PricePerDay:    pricePerDay,
				ReservedDate:   reservedDate,
				PickUpTime:     pickUpTime,
				DropOfTime:     dropOff,
				CarId:          carID,
				LocationId:     locationId,
				CustomerId:     customer.Uid,
			}
			if DB.Where("id = ?", car.Id).First(&models.Car{}).RowsAffected == 0 {
				DB.Create(&car)
			} else {
				DB.Model(&car).Updates(car)
			}
			if DB.Where("id = ?", location.Id).First(&models.Location{}).RowsAffected == 0 {
				DB.Create(&location)
			} else {
				DB.Model(&location).Updates(location)
			}
			if DB.Where("uid = ?", customer.Uid).First(&models.Customer{}).RowsAffected == 0 {
				DB.Create(&customer)
			} else {
				DB.Model(&customer).Updates(customer)
			}
			records = append(records, bookRecord)

		}
	}
	for _, record := range records {
		if DB.Where("reservation_num = ?", record.ReservationNum).First(&models.BookRecord{}).RowsAffected == 0 {
			DB.Create(&record)
		} else {
			DB.Model(&record).Updates(record)
		}
	}
}
