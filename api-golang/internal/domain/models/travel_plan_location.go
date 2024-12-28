package models

type TravelPlanLocation struct {
	ID           uint       `gorm:"primaryKey" json:"id"`
	TravelPlanID uint       `gorm:"not null" json:"travel_plan_id"`
	TravelPlan   TravelPlan `gorm:"foreignKey:TravelPlanID;constraint:onDelete:CASCADE" json:"travel_plan"`
	LocationID   uint       `gorm:"not null" json:"location_id"`
	Location     Location   `gorm:"foreignKey:LocationID;constraint:onDelete:CASCADE" json:"location"`
	VisitOrder   int        `gorm:"not null" json:"visit_order"`
}
