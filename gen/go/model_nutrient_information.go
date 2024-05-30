/*
 * Japanese Food Nutrient API
 *
 * API to access nutrient information of Japanese food ingredients
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type NutrientInformation struct {

	// Name of the nutrient
	Name string `json:"name,omitempty"`

	// Amount of the nutrient per 100g of the food item
	Amount float32 `json:"amount,omitempty"`

	// Unit of measurement for the nutrient amount
	Unit string `json:"unit,omitempty"`
}
