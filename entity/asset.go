package entity

import (
	"time"

	"github.com/satori/uuid"
)

// Asset asset entity
type Asset struct {
	ID string `json:"id" xorm:"id"`

	ProductID         string    `json:"product_id" xorm:"product_id"`
	SerialNumber      string    `json:"serial_number" xorm:"serial_number"`
	Status            int       `json:"status" xorm:"status"`
	PurchaseDate      time.Time `json:"purchase_date" xorm:"purchase_date"`
	PurchasePrice     float32   `json:"purchase_price" xorm:"purchase_price"`
	SupplierCompanyID string    `json:"supplier_company_id" xorm:"supplier_company_id"`
	SalvageValue      float32   `json:"salvage_value" xorm:"salvage_value"`

	TerminalID                             *string           `json:"terminal_id" xorm:"terminal_id"`
	LocationName                           *string           `json:"location_name" xorm:"location_name"`
	LocationIdentity                       *LocationIdentity `json:"location_identity" xorm:"location_identity"`
	Address                                *string           `json:"address" xorm:"address"`
	AdministrativeArea                     *string           `json:"administrative_area" xorm:"administrative_area"`
	Latitude                               *string           `json:"latitude" xorm:"latitude"`
	Longitude                              *string           `json:"longitude" xorm:"longitude"`
	Altitude                               *string           `json:"altitude" xorm:"altitude"`
	LocationTimeAndAccess                  *string           `json:"location_time_and_access" xorm:"location_time_and_access"`
	PICLocation                            *string           `json:"pic_location" xorm:"pic_location"`
	PICArea                                *string           `json:"pic_area" xorm:"pic_area"`
	PhoneNumber                            *string           `json:"phone_number" xorm:"phone_number"`
	Day                                    *string           `json:"day" xorm:"day"`
	Hour                                   *string           `json:"hour" xorm:"hour"`
	MainEnergySource                       *string           `json:"main_energy_source" xorm:"main_energy_source"`
	AlternativeEnergySource                *string           `json:"alternative_energy_source" xorm:"alternative_energy_source"`
	ExtraInformation                       *string           `json:"extra_information" xorm:"extra_information"`
	MinimumVoltage                         *string           `json:"minimum_voltage" xorm:"minimum_voltage"`
	MaximumVoltage                         *string           `json:"maximum_voltage" xorm:"maximum_voltage"`
	GroundingNtoG                          *string           `json:"grounding_n_to_g" xorm:"grounding_n_to_g"`
	ElectricalEquipmentAvailableInLocation *string           `json:"electrical_equipment_available_in_location" xorm:"electrical_equipment_available_in_location"`
	LCProvide                              *string           `json:"lc_provide" xorm:"lc_provide"`
	SatelitteName                          *string           `json:"satelitte_name" xorm:"satelitte_name"`
	SpotBeam                               *string           `json:"spot_beam" xorm:"spot_beam"`
	OperationBand                          *OperationBand    `json:"operation_band" xorm:"operation_band"`
	VSATSystem                             *VSAT             `json:"vsat_system" xorm:"vsat_system"`
	ModemBrand                             *string           `json:"modem_brand" xorm:"modem_brand"`
	ModemType                              *string           `json:"modem_type" xorm:"modem_type"`
	ModemSN                                *string           `json:"modem_sn" xorm:"modem_sn"`
	AntennaDiameter                        *string           `json:"antenna_diameter" xorm:"antenna_diameter"`
	AntennaBrand                           *string           `json:"antenna_brand" xorm:"antenna_brand"`
	AntennaType                            *string           `json:"antenna_type" xorm:"antenna_type"`
	AntennaSN                              *string           `json:"antenna_sn" xorm:"antenna_sn"`
	SspaBuc                                *string           `json:"sspa_buc" xorm:"sspa_buc"`
	SspaBucBrand                           *string           `json:"sspa_buc_brand" xorm:"sspa_buc_brand"`
	SspaBucType                            *string           `json:"sspa_buc_type" xorm:"sspa_buc_type"`
	SspaBucSN                              *string           `json:"sspa_buc_sn" xorm:"sspa_buc_sn"`
	LnBBrand                               *string           `json:"l_n_b_brand" xorm:"l_n_b_brand"`
	LnBType                                *string           `json:"l_n_b_type" xorm:"l_n_b_type"`
	LnBSN                                  *string           `json:"l_n_b_sn" xorm:"l_n_b_sn"`
	SwitchType                             *string           `json:"switch_type" xorm:"switch_type"`
	SwitchBrand                            *string           `json:"switch_brand" xorm:"switch_brand" `
	SwitchSN                               *string           `json:"switch_sn" xorm:"switch_sn"`
	IPManagementModem                      *string           `json:"ipmanagement_modem" xorm:"ipmanagement_modem"`
	IPModem                                *string           `json:"ipmodem" xorm:"ipmodem"`
	IPRouter                               *string           `json:"iprouter" xorm:"iprouter"`
	IPAP1                                  *string           `json:"ipap_1" xorm:"ipap_1"`
	IPAP2                                  *string           `json:"ipap_2" xorm:"ipap_2"`
	UserType                               *AssetUserType    `json:"user_type" xorm:"user_type"`
	AntennaLocation                        *string           `json:"antenna_location" xorm:"antenna_location"`
	AntennaLocationType                    *string           `json:"antenna_location_type" xorm:"antenna_location_type"`
	AntennaPedestalType                    *string           `json:"antenna_pedestal_type" xorm:"antenna_pedestal_type"`
	IndoorUnitLocation                     *string           `json:"indoor_unit_location" xorm:"indoor_unit_location"`
	IndoorUnit                             *string           `json:"indoor_unit" xorm:"indoor_unit"`
	SupportingParts                        *string           `json:"supporting_parts" xorm:"supporting_parts"`
	OutdoorAP1Location                     *string           `json:"outdoor_ap1location" xorm:"outdoor_ap1location"`
	OutdoorAP1Type                         *string           `json:"outdoor_ap_1_type" xorm:"outdoor_ap_1_type"`
	OutdoorAP1Brand                        *string           `json:"outdoor_ap_1_brand" xorm:"outdoor_ap_1_brand"`
	OutdoorAP1SN                           *string           `json:"outdoor_ap_1_sn" xorm:"outdoor_ap_1_sn"`
	OutdoorAP1Range                        *string           `json:"outdoor_ap_1_range" xorm:"outdoor_ap_1_range"`
	OutdoorAP1Audience                     *string           `json:"outdoor_ap_1_audience" xorm:"outdoor_ap_1_audience"`
	OutdoorAP1Security                     *string           `json:"outdoor_ap_1_security" xorm:"outdoor_ap_1_security"`
	OutdoorAP1DurationSettingsBandwidthCos *string           `json:"outdoor_ap_1_duration_settings_bandwidth_cos" xorm:"outdoor_ap_1_duration_settings_bandwidth_cos"`
	OutdoorAP2Location                     *string           `json:"outdoor_ap_2_location" xorm:"outdoor_ap2location"`
	OutdoorAP2Type                         *string           `json:"outdoor_ap_2_type" xorm:"outdoor_ap2type"`
	OutdoorAP2Brand                        *string           `json:"outdoor_ap_2_brand" xorm:"outdoor_ap2brand"`
	OutdoorAP2SN                           *string           `json:"outdoor_ap_2_sn" xorm:"outdoor_ap2sn"`
	OutdoorAP2Range                        *string           `json:"outdoor_ap_2_range" xorm:"outdoor_ap2range"`
	OutdoorAP2Audience                     *string           `json:"outdoor_ap_2_audience" xorm:"outdoor_ap2audience"`
	OutdoorAP2Security                     *string           `json:"outdoor_ap_2_security" xorm:"outdoor_ap2security"`
	OutdoorAP2DurationSettingsBandwidthCos *string           `json:"outdoor_ap_2_duration_settings_bandwidth_cos" xorm:"outdoor_ap2duration_settings_bandwidth_cos"`
	PowerStabilizerBrand                   *string           `json:"power_stabilizer_brand" xorm:"power_stabilizer_brand"`
	PowerStabilizerType                    *string           `json:"power_stabilizer_type" xorm:"power_stabilizer_type"`
	PowerStabilizerSN                      *string           `json:"power_stabilizer_sn" xorm:"power_stabilizer_sn"`
	RackBrand                              *string           `json:"rack_brand" xorm:"rack_brand"`
	RackType                               *string           `json:"rack_type" xorm:"rack_type"`
	RackSN                                 *string           `json:"rack_sn" xorm:"rack_sn"`
	Notes                                  *string           `json:"notes" xorm:"notes"`

	CreatedBy *string    `json:"created_by" xorm:"created_by"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

// AssetUserType :nodoc
type AssetUserType = int32

// LocationIdentity :nodoc
type LocationIdentity = int32

// OperationBand :nodoc
type OperationBand = int32

// VSAT :nodoc
type VSAT = int32

const (
	// TDMA :nodoc
	TDMA = iota + 1
	// SCPC :nodoc
	SCPC
	// DSCPC :nodoc
	DSCPC
	// MCPC :nodoc
	MCPC
	// MXDMA :nodoc
	MXDMA
)

const (
	// KuBand :nodoc
	KuBand = iota + 1
	// CBand :nodoc
	CBand
)

const (
	// UserSchool :nodoc
	UserSchool AssetUserType = iota
	// UserOffice :nodoc
	UserOffice
	// UserGovernment :nodoc
	UserGovernment
	// UserHealthcare :nodoc
	UserHealthcare
	// UserPublicService :nodoc
	UserPublicService
	// UserBalaiLatihanKerja :nodoc
	UserBalaiLatihanKerja
	// UserTourism :nodoc
	UserTourism
	// UserReligion :nodoc
	UserReligion
	// UserKarantinaPertanian :nodoc
	UserKarantinaPertanian
	// UserAirport :nodoc
	UserAirport
	// UserPLBN :nodoc
	UserPLBN
)

const (
	// LocationSchool :nodoc
	LocationSchool LocationIdentity = iota
	// LocationOffice :nodoc
	LocationOffice
	// LocationGovernment :nodoc
	LocationGovernment
	// LocationHealthcare :nodoc
	LocationHealthcare
	// LocationPublicService :nodoc
	LocationPublicService
	// LocationBalaiLatihanKerja :nodoc
	LocationBalaiLatihanKerja
	// LocationTourism :nodoc
	LocationTourism
	// LocationReligion :nodoc
	LocationReligion
	// LocationKarantinaPertanian :nodoc
	LocationKarantinaPertanian
	// LocationAirport :nodoc
	LocationAirport
	// LocationPLBN :nodoc
	LocationPLBN
)

// AssetRequest :nodoc
type AssetRequest struct {
	ID string `json:"id" xorm:"id"`

	ProductID         string    `json:"product_id" xorm:"product_id"`
	SerialNumber      string    `json:"serial_number" xorm:"serial_number"`
	Status            int       `json:"status" xorm:"status"`
	PurchaseDate      time.Time `json:"purchase_date" xorm:"purchase_date"`
	PurchasePrice     float32   `json:"purchase_price" xorm:"purchase_price"`
	SupplierCompanyID string    `json:"supplier_company_id" xorm:"supplier_company_id"`
	SalvageValue      float32   `json:"salvage_value" xorm:"salvage_value"`
	CreatedBy         *string   `json:"created_by" xorm:"created_by"`

	TerminalID                             *string           `json:"terminal_id" xorm:"terminal_id"`
	LocationName                           *string           `json:"location_name" xorm:"location_name"`
	LocationIdentity                       *LocationIdentity `json:"location_identity" xorm:"location_identity"`
	Address                                *string           `json:"address" xorm:"address"`
	AdministrativeArea                     *string           `json:"administrative_area" xorm:"administrative_area"`
	Latitude                               *string           `json:"latitude" xorm:"latitude"`
	Longitude                              *string           `json:"longitude" xorm:"longitude"`
	Altitude                               *string           `json:"altitude" xorm:"altitude"`
	LocationTimeAndAccess                  *string           `json:"location_time_and_access" xorm:"location_time_and_access"`
	PICLocation                            *string           `json:"pic_location" xorm:"pic_location"`
	PICArea                                *string           `json:"pic_area" xorm:"pic_area"`
	PhoneNumber                            *string           `json:"phone_number" xorm:"phone_number"`
	Day                                    *string           `json:"day" xorm:"day"`
	Hour                                   *string           `json:"hour" xorm:"hour"`
	MainEnergySource                       *string           `json:"main_energy_source" xorm:"main_energy_source"`
	AlternativeEnergySource                *string           `json:"alternative_energy_source" xorm:"alternative_energy_source"`
	ExtraInformation                       *string           `json:"extra_information" xorm:"extra_information"`
	MinimumVoltage                         *string           `json:"minimum_voltage" xorm:"minimum_voltage"`
	MaximumVoltage                         *string           `json:"maximum_voltage" xorm:"maximum_voltage"`
	GroundingNtoG                          *string           `json:"grounding_n_to_g" xorm:"grounding_n_to_g"`
	ElectricalEquipmentAvailableInLocation *string           `json:"electrical_equipment_available_in_location" xorm:"electrical_equipment_available_in_location"`
	LCProvide                              *string           `json:"lc_provide" xorm:"lc_provide"`
	SatelitteName                          *string           `json:"satelitte_name" xorm:"satelitte_name"`
	SpotBeam                               *string           `json:"spot_beam" xorm:"spot_beam"`
	OperationBand                          *OperationBand    `json:"operation_band" xorm:"operation_band"`
	VSATSystem                             *VSAT             `json:"vsat_system" xorm:"vsat_system"`
	ModemBrand                             *string           `json:"modem_brand" xorm:"modem_brand"`
	ModemType                              *string           `json:"modem_type" xorm:"modem_type"`
	ModemSN                                *string           `json:"modem_sn" xorm:"modem_sn"`
	AntennaDiameter                        *string           `json:"antenna_diameter" xorm:"antenna_diameter"`
	AntennaBrand                           *string           `json:"antenna_brand" xorm:"antenna_brand"`
	AntennaType                            *string           `json:"antenna_type" xorm:"antenna_type"`
	AntennaSN                              *string           `json:"antenna_sn" xorm:"antenna_sn"`
	SspaBuc                                *string           `json:"sspa_buc" xorm:"sspa_buc"`
	SspaBucBrand                           *string           `json:"sspa_buc_brand" xorm:"sspa_buc_brand"`
	SspaBucType                            *string           `json:"sspa_buc_type" xorm:"sspa_buc_type"`
	SspaBucSN                              *string           `json:"sspa_buc_sn" xorm:"sspa_buc_sn"`
	LnBBrand                               *string           `json:"l_n_b_brand" xorm:"l_n_b_brand"`
	LnBType                                *string           `json:"l_n_b_type" xorm:"l_n_b_type"`
	LnBSN                                  *string           `json:"l_n_b_sn" xorm:"l_n_b_sn"`
	SwitchType                             *string           `json:"switch_type" xorm:"switch_type"`
	SwitchBrand                            *string           `json:"switch_brand" xorm:"switch_brand" `
	SwitchSN                               *string           `json:"switch_sn" xorm:"switch_sn"`
	IPManagementModem                      *string           `json:"ipmanagement_modem" xorm:"ipmanagement_modem"`
	IPModem                                *string           `json:"ipmodem" xorm:"ipmodem"`
	IPRouter                               *string           `json:"iprouter" xorm:"iprouter"`
	IPAP1                                  *string           `json:"ipap_1" xorm:"ipap_1"`
	IPAP2                                  *string           `json:"ipap_2" xorm:"ipap_2"`
	UserType                               *AssetUserType    `json:"user_type" xorm:"user_type"`
	AntennaLocation                        *string           `json:"antenna_location" xorm:"antenna_location"`
	AntennaLocationType                    *string           `json:"antenna_location_type" xorm:"antenna_location_type"`
	AntennaPedestalType                    *string           `json:"antenna_pedestal_type" xorm:"antenna_pedestal_type"`
	IndoorUnitLocation                     *string           `json:"indoor_unit_location" xorm:"indoor_unit_location"`
	IndoorUnit                             *string           `json:"indoor_unit" xorm:"indoor_unit"`
	SupportingParts                        *string           `json:"supporting_parts" xorm:"supporting_parts"`
	OutdoorAP1Location                     *string           `json:"outdoor_ap1location" xorm:"outdoor_ap1location"`
	OutdoorAP1Type                         *string           `json:"outdoor_ap_1_type" xorm:"outdoor_ap_1_type"`
	OutdoorAP1Brand                        *string           `json:"outdoor_ap_1_brand" xorm:"outdoor_ap_1_brand"`
	OutdoorAP1SN                           *string           `json:"outdoor_ap_1_sn" xorm:"outdoor_ap_1_sn"`
	OutdoorAP1Range                        *string           `json:"outdoor_ap_1_range" xorm:"outdoor_ap_1_range"`
	OutdoorAP1Audience                     *string           `json:"outdoor_ap_1_audience" xorm:"outdoor_ap_1_audience"`
	OutdoorAP1Security                     *string           `json:"outdoor_ap_1_security" xorm:"outdoor_ap_1_security"`
	OutdoorAP1DurationSettingsBandwidthCos *string           `json:"outdoor_ap_1_duration_settings_bandwidth_cos" xorm:"outdoor_ap_1_duration_settings_bandwidth_cos"`
	OutdoorAP2Location                     *string           `json:"outdoor_ap_2_location" xorm:"outdoor_ap2location"`
	OutdoorAP2Type                         *string           `json:"outdoor_ap_2_type" xorm:"outdoor_ap2type"`
	OutdoorAP2Brand                        *string           `json:"outdoor_ap_2_brand" xorm:"outdoor_ap2brand"`
	OutdoorAP2SN                           *string           `json:"outdoor_ap_2_sn" xorm:"outdoor_ap2sn"`
	OutdoorAP2Range                        *string           `json:"outdoor_ap_2_range" xorm:"outdoor_ap2range"`
	OutdoorAP2Audience                     *string           `json:"outdoor_ap_2_audience" xorm:"outdoor_ap2audience"`
	OutdoorAP2Security                     *string           `json:"outdoor_ap_2_security" xorm:"outdoor_ap2security"`
	OutdoorAP2DurationSettingsBandwidthCos *string           `json:"outdoor_ap_2_duration_settings_bandwidth_cos" xorm:"outdoor_ap2duration_settings_bandwidth_cos"`
	PowerStabilizerBrand                   *string           `json:"power_stabilizer_brand" xorm:"power_stabilizer_brand"`
	PowerStabilizerType                    *string           `json:"power_stabilizer_type" xorm:"power_stabilizer_type"`
	PowerStabilizerSN                      *string           `json:"power_stabilizer_sn" xorm:"power_stabilizer_sn"`
	RackBrand                              *string           `json:"rack_brand" xorm:"rack_brand"`
	RackType                               *string           `json:"rack_type" xorm:"rack_type"`
	RackSN                                 *string           `json:"rack_sn" xorm:"rack_sn"`
	Notes                                  *string           `json:"notes" xorm:"notes"`
}

// AssetGroup asset with mapped data
type AssetGroup struct {
	Asset
	CurrentValuation float32      `json:"current_valuation"`
	Product          ProductGroup `json:"product"`
	Warehouse        []Warehouse  `json:"warehouses"`
	Company          CompanyGroup `json:"company"`
	CreatedByUser    *User        `json:"created_by_user"`
}

// AssetChangeSet change set forasset
type AssetChangeSet struct {
	ProductID         string    `json:"-" xorm:"product_id"`
	SerialNumber      string    `json:"serial_number" xorm:"serial_number"`
	Status            int       `json:"status" xorm:"status"`
	PurchaseDate      time.Time `json:"purchase_date" xorm:"purchase_date"`
	PurchasePrice     float32   `json:"purchase_price" xorm:"purchase_price"`
	SupplierCompanyID string    `json:"-" xorm:"supplier_company_id"`
	SalvageValue      float32   `json:"salvage_value" xorm:"salvage_value"`

	TerminalID                             *string           `json:"terminal_id" xorm:"terminal_id"`
	LocationName                           *string           `json:"location_name" xorm:"location_name"`
	LocationIdentity                       *LocationIdentity `json:"location_identity" xorm:"location_identity"`
	Address                                *string           `json:"address" xorm:"address"`
	AdministrativeArea                     *string           `json:"administrative_area" xorm:"administrative_area"`
	Latitude                               *string           `json:"latitude" xorm:"latitude"`
	Longitude                              *string           `json:"longitude" xorm:"longitude"`
	Altitude                               *string           `json:"altitude" xorm:"altitude"`
	LocationTimeAndAccess                  *string           `json:"location_time_and_access" xorm:"location_time_and_access"`
	PICLocation                            *string           `json:"pic_location" xorm:"pic_location"`
	PICArea                                *string           `json:"pic_area" xorm:"pic_area"`
	PhoneNumber                            *string           `json:"phone_number" xorm:"phone_number"`
	Day                                    *string           `json:"day" xorm:"day"`
	Hour                                   *string           `json:"hour" xorm:"hour"`
	MainEnergySource                       *string           `json:"main_energy_source" xorm:"main_energy_source"`
	AlternativeEnergySource                *string           `json:"alternative_energy_source" xorm:"alternative_energy_source"`
	ExtraInformation                       *string           `json:"extra_information" xorm:"extra_information"`
	MinimumVoltage                         *string           `json:"minimum_voltage" xorm:"minimum_voltage"`
	MaximumVoltage                         *string           `json:"maximum_voltage" xorm:"maximum_voltage"`
	GroundingNtoG                          *string           `json:"grounding_n_to_g" xorm:"grounding_n_to_g"`
	ElectricalEquipmentAvailableInLocation *string           `json:"electrical_equipment_available_in_location" xorm:"electrical_equipment_available_in_location"`
	LCProvide                              *string           `json:"lc_provide" xorm:"lc_provide"`
	SatelitteName                          *string           `json:"satelitte_name" xorm:"satelitte_name"`
	SpotBeam                               *string           `json:"spot_beam" xorm:"spot_beam"`
	OperationBand                          *OperationBand    `json:"operation_band" xorm:"operation_band"`
	VSATSystem                             *VSAT             `json:"vsat_system" xorm:"vsat_system"`
	ModemBrand                             *string           `json:"modem_brand" xorm:"modem_brand"`
	ModemType                              *string           `json:"modem_type" xorm:"modem_type"`
	ModemSN                                *string           `json:"modem_sn" xorm:"modem_sn"`
	AntennaDiameter                        *string           `json:"antenna_diameter" xorm:"antenna_diameter"`
	AntennaBrand                           *string           `json:"antenna_brand" xorm:"antenna_brand"`
	AntennaType                            *string           `json:"antenna_type" xorm:"antenna_type"`
	AntennaSN                              *string           `json:"antenna_sn" xorm:"antenna_sn"`
	SspaBuc                                *string           `json:"sspa_buc" xorm:"sspa_buc"`
	SspaBucBrand                           *string           `json:"sspa_buc_brand" xorm:"sspa_buc_brand"`
	SspaBucType                            *string           `json:"sspa_buc_type" xorm:"sspa_buc_type"`
	SspaBucSN                              *string           `json:"sspa_buc_sn" xorm:"sspa_buc_sn"`
	LnBBrand                               *string           `json:"l_n_b_brand" xorm:"l_n_b_brand"`
	LnBType                                *string           `json:"l_n_b_type" xorm:"l_n_b_type"`
	LnBSN                                  *string           `json:"l_n_b_sn" xorm:"l_n_b_sn"`
	SwitchType                             *string           `json:"switch_type" xorm:"switch_type"`
	SwitchBrand                            *string           `json:"switch_brand" xorm:"switch_brand" `
	SwitchSN                               *string           `json:"switch_sn" xorm:"switch_sn"`
	IPManagementModem                      *string           `json:"ipmanagement_modem" xorm:"ipmanagement_modem"`
	IPModem                                *string           `json:"ipmodem" xorm:"ipmodem"`
	IPRouter                               *string           `json:"iprouter" xorm:"iprouter"`
	IPAP1                                  *string           `json:"ipap_1" xorm:"ipap_1"`
	IPAP2                                  *string           `json:"ipap_2" xorm:"ipap_2"`
	UserType                               *AssetUserType    `json:"user_type" xorm:"user_type"`
	AntennaLocation                        *string           `json:"antenna_location" xorm:"antenna_location"`
	AntennaLocationType                    *string           `json:"antenna_location_type" xorm:"antenna_location_type"`
	AntennaPedestalType                    *string           `json:"antenna_pedestal_type" xorm:"antenna_pedestal_type"`
	IndoorUnitLocation                     *string           `json:"indoor_unit_location" xorm:"indoor_unit_location"`
	IndoorUnit                             *string           `json:"indoor_unit" xorm:"indoor_unit"`
	SupportingParts                        *string           `json:"supporting_parts" xorm:"supporting_parts"`
	OutdoorAP1Location                     *string           `json:"outdoor_ap1location" xorm:"outdoor_ap1location"`
	OutdoorAP1Type                         *string           `json:"outdoor_ap_1_type" xorm:"outdoor_ap_1_type"`
	OutdoorAP1Brand                        *string           `json:"outdoor_ap_1_brand" xorm:"outdoor_ap_1_brand"`
	OutdoorAP1SN                           *string           `json:"outdoor_ap_1_sn" xorm:"outdoor_ap_1_sn"`
	OutdoorAP1Range                        *string           `json:"outdoor_ap_1_range" xorm:"outdoor_ap_1_range"`
	OutdoorAP1Audience                     *string           `json:"outdoor_ap_1_audience" xorm:"outdoor_ap_1_audience"`
	OutdoorAP1Security                     *string           `json:"outdoor_ap_1_security" xorm:"outdoor_ap_1_security"`
	OutdoorAP1DurationSettingsBandwidthCos *string           `json:"outdoor_ap_1_duration_settings_bandwidth_cos" xorm:"outdoor_ap_1_duration_settings_bandwidth_cos"`
	OutdoorAP2Location                     *string           `json:"outdoor_ap_2_location" xorm:"outdoor_ap2location"`
	OutdoorAP2Type                         *string           `json:"outdoor_ap_2_type" xorm:"outdoor_ap2type"`
	OutdoorAP2Brand                        *string           `json:"outdoor_ap_2_brand" xorm:"outdoor_ap2brand"`
	OutdoorAP2SN                           *string           `json:"outdoor_ap_2_sn" xorm:"outdoor_ap2sn"`
	OutdoorAP2Range                        *string           `json:"outdoor_ap_2_range" xorm:"outdoor_ap2range"`
	OutdoorAP2Audience                     *string           `json:"outdoor_ap_2_audience" xorm:"outdoor_ap2audience"`
	OutdoorAP2Security                     *string           `json:"outdoor_ap_2_security" xorm:"outdoor_ap2security"`
	OutdoorAP2DurationSettingsBandwidthCos *string           `json:"outdoor_ap_2_duration_settings_bandwidth_cos" xorm:"outdoor_ap2duration_settings_bandwidth_cos"`
	PowerStabilizerBrand                   *string           `json:"power_stabilizer_brand" xorm:"power_stabilizer_brand"`
	PowerStabilizerType                    *string           `json:"power_stabilizer_type" xorm:"power_stabilizer_type"`
	PowerStabilizerSN                      *string           `json:"power_stabilizer_sn" xorm:"power_stabilizer_sn"`
	RackBrand                              *string           `json:"rack_brand" xorm:"rack_brand"`
	RackType                               *string           `json:"rack_type" xorm:"rack_type"`
	RackSN                                 *string           `json:"rack_sn" xorm:"rack_sn"`
	Notes                                  *string           `json:"notes" xorm:"notes"`
}

// NewAsset create newasset
func NewAsset(
	request AssetRequest,
) (asset Asset, err error) {
	asset = Asset{
		ID:                                     uuid.NewV4().String(),
		ProductID:                              request.ProductID,
		SerialNumber:                           request.SerialNumber,
		Status:                                 request.Status,
		PurchaseDate:                           request.PurchaseDate,
		PurchasePrice:                          request.PurchasePrice,
		SupplierCompanyID:                      request.SupplierCompanyID,
		SalvageValue:                           request.SalvageValue,
		CreatedBy:                              request.CreatedBy,
		TerminalID:                             request.TerminalID,
		LocationName:                           request.LocationName,
		LocationIdentity:                       request.LocationIdentity,
		Address:                                request.Address,
		AdministrativeArea:                     request.AdministrativeArea,
		Latitude:                               request.Latitude,
		Longitude:                              request.Longitude,
		Altitude:                               request.Altitude,
		LocationTimeAndAccess:                  request.LocationTimeAndAccess,
		PICLocation:                            request.PICLocation,
		PICArea:                                request.PICArea,
		PhoneNumber:                            request.PhoneNumber,
		Day:                                    request.Day,
		Hour:                                   request.Hour,
		MainEnergySource:                       request.MainEnergySource,
		AlternativeEnergySource:                request.AlternativeEnergySource,
		ExtraInformation:                       request.ExtraInformation,
		MinimumVoltage:                         request.MinimumVoltage,
		MaximumVoltage:                         request.MaximumVoltage,
		GroundingNtoG:                          request.GroundingNtoG,
		ElectricalEquipmentAvailableInLocation: request.ElectricalEquipmentAvailableInLocation,
		LCProvide:                              request.LCProvide,
		SatelitteName:                          request.SatelitteName,
		SpotBeam:                               request.SpotBeam,
		OperationBand:                          request.OperationBand,
		VSATSystem:                             request.VSATSystem,
		ModemBrand:                             request.ModemBrand,
		ModemType:                              request.ModemType,
		ModemSN:                                request.ModemSN,
		AntennaDiameter:                        request.AntennaDiameter,
		AntennaBrand:                           request.AntennaBrand,
		AntennaType:                            request.AntennaType,
		AntennaSN:                              request.AntennaSN,
		SspaBuc:                                request.SspaBuc,
		SspaBucBrand:                           request.SspaBucBrand,
		SspaBucType:                            request.SspaBucType,
		SspaBucSN:                              request.SspaBucSN,
		LnBBrand:                               request.LnBBrand,
		LnBType:                                request.LnBType,
		LnBSN:                                  request.LnBSN,
		SwitchType:                             request.SwitchType,
		SwitchBrand:                            request.SwitchBrand,
		SwitchSN:                               request.SwitchSN,
		IPManagementModem:                      request.IPManagementModem,
		IPModem:                                request.IPModem,
		IPRouter:                               request.IPRouter,
		IPAP1:                                  request.IPAP1,
		IPAP2:                                  request.IPAP2,
		UserType:                               request.UserType,
		AntennaLocation:                        request.AntennaLocation,
		AntennaLocationType:                    request.AntennaLocationType,
		AntennaPedestalType:                    request.AntennaPedestalType,
		IndoorUnitLocation:                     request.IndoorUnitLocation,
		IndoorUnit:                             request.IndoorUnit,
		SupportingParts:                        request.SupportingParts,
		OutdoorAP1Location:                     request.OutdoorAP1Location,
		OutdoorAP1Type:                         request.OutdoorAP1Type,
		OutdoorAP1Brand:                        request.OutdoorAP1Brand,
		OutdoorAP1SN:                           request.OutdoorAP1SN,
		OutdoorAP1Range:                        request.OutdoorAP1Range,
		OutdoorAP1Audience:                     request.OutdoorAP1Audience,
		OutdoorAP1Security:                     request.OutdoorAP1Security,
		OutdoorAP1DurationSettingsBandwidthCos: request.OutdoorAP1DurationSettingsBandwidthCos,
		OutdoorAP2Location:                     request.OutdoorAP2Location,
		OutdoorAP2Type:                         request.OutdoorAP2Type,
		OutdoorAP2Brand:                        request.OutdoorAP2Brand,
		OutdoorAP2SN:                           request.OutdoorAP2SN,
		OutdoorAP2Range:                        request.OutdoorAP2Range,
		OutdoorAP2Audience:                     request.OutdoorAP2Audience,
		OutdoorAP2Security:                     request.OutdoorAP2Security,
		OutdoorAP2DurationSettingsBandwidthCos: request.OutdoorAP2DurationSettingsBandwidthCos,
		PowerStabilizerBrand:                   request.PowerStabilizerBrand,
		PowerStabilizerType:                    request.PowerStabilizerType,
		PowerStabilizerSN:                      request.PowerStabilizerSN,
		RackBrand:                              request.RackBrand,
		RackType:                               request.RackType,
		RackSN:                                 request.RackSN,
		Notes:                                  request.Notes,
	}
	return
}
