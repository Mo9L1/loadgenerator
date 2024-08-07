package behaviors

import (
	"fmt"
	"github.com/go-faker/faker/v4"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// generateVerifyCode generates a 6-digit verification code consisting of letters and numbers.
func generateVerifyCode() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const length = 6
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

	code := make([]byte, length)
	for i := range code {
		code[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(code)
}

func GenerateTripId() string {
	// 设置随机数种子
	rand.Seed(time.Now().UnixNano())

	// 定义可能的开头字母
	letters := []rune{'Z', 'T', 'K', 'G', 'D'}

	// 随机选择一个字母
	startLetter := letters[rand.Intn(len(letters))]

	// 生成三个随机数字
	randomNumber := rand.Intn(1000)

	// 格式化成三位数字，不足三位前面补零
	MockedTripID := fmt.Sprintf("%c%03d", startLetter, randomNumber)

	return MockedTripID
}

// toLowerCaseAndRemoveSpaces converts a given string to all lower case
// and removes all spaces.
func toLowerCaseAndRemoveSpaces(input string) string {
	lowercased := strings.ToLower(input)
	noSpaces := strings.ReplaceAll(lowercased, " ", "")
	return noSpaces
}

// generateRandomFood generates a random food item from a predefined list of 50 kinds of food.
func generateRandomFood() string {
	// Predefined list of food items
	foodList := []string{
		"Pizza", "Burger", "Pasta", "Sushi", "Tacos", "Salad", "Steak", "Soup", "Sandwich", "Fries",
		"Ice Cream", "Cake", "Donut", "Chocolate", "Apple", "Banana", "Orange", "Grapes", "Strawberry", "Blueberry",
		"Mango", "Pineapple", "Watermelon", "Kiwi", "Avocado", "Tomato", "Cucumber", "Carrot", "Broccoli", "Spinach",
		"Chicken", "Beef", "Pork", "Lamb", "Fish", "Shrimp", "Crab", "Lobster", "Eggs", "Cheese",
		"Yogurt", "Milk", "Butter", "Bread", "Rice", "Pasta", "Noodles", "Cereal", "Oatmeal", "Honey",
	}

	// Seed the random number generator to ensure different results each run
	rand.Seed(time.Now().UnixNano())

	// Generate a random index to pick a food item
	randomIndex := rand.Intn(len(foodList))

	// Return the randomly selected food item
	return foodList[randomIndex]
}

// generateRandomStoreName generates a random store name from a predefined list of 30 kinds of store names.
func generateRandomStoreName() string {
	// Predefined list of store names
	storeNames := []string{
		"Grocery Mart", "Tech World", "Fashion Hub", "Book Haven", "Toy Land",
		"Pet Paradise", "Home Essentials", "Beauty Bliss", "Sports Store", "Gadget Garage",
		"Furniture Factory", "Shoe Stop", "Pharmacy Plus", "Hardware Haven", "Electronics Emporium",
		"Music Mania", "Garden Goods", "Office Outlet", "Auto Accessories", "Craft Corner",
		"Gift Gallery", "Jewelry Junction", "Bakery Bliss", "Café Corner", "Fitness Freak",
		"Outdoor Outfitters", "Travel Treasures", "Kids' Kingdom", "Vintage Vault", "Wine World",
	}

	// Seed the random number generator to ensure different results each run
	rand.Seed(time.Now().UnixNano())

	// Generate a random index to pick a store name
	randomIndex := rand.Intn(len(storeNames))

	// Return the randomly selected store name
	return storeNames[randomIndex]
}

// helper function for Order Service
// RandomDecimalStringBetween 生成并返回两个整数之间的一位小数形式的随机数字符串，包括边界值。
func RandomDecimalStringBetween(min, max int) string {
	rand.Seed(time.Now().UnixNano())
	randomInt := rand.Intn(max-min+1) + min              // 生成[min, max]范围内的随机整数
	decimalValue := float64(randomInt) * 0.1             // 将整数转换为一位小数
	return strconv.FormatFloat(decimalValue, 'f', 1, 64) // 转换为一位小数的字符串形式
}

// RandomProvincialCapitalEN 随机返回一个中国省会城市的英文名称
func RandomProvincialCapitalEN() string {
	rand.Seed(time.Now().UnixNano())
	return provincialCapitalsEN[rand.Intn(len(provincialCapitalsEN))]
}

// 中国省会城市的英文列表
var provincialCapitalsEN = []string{
	"Beijing", "Shanghai", "Tianjin", "Chongqing",
	"Shijiazhuang", "Taiyuan", "Hohhot", "Shenyang", "Changchun", "Harbin",
	"Nanjing", "Hangzhou", "Hefei", "Fuzhou", "Nanchang", "Jinan", "Zhengzhou", "Wuhan", "Changsha", "Guangzhou",
	"Nanning", "Haikou", "Chengdu", "Guiyang", "Kunming", "Lhasa", "Xi'an", "Lanzhou", "Xining", "Yinchuan",
	"Urumqi", "Taipei",
}

// GetTrainTicketClass 随机返回高铁票等级。
// 有5%的概率返回"FirstClass"（头等座），
// 15%的概率返回"BusinessClass"（一等座），
// 剩余80%的概率返回"EconomyClass"（二等座）。
func GetTrainTicketClass() int {
	rand.Seed(time.Now().UnixNano()) // 确保每次运行时随机数种子不同

	probability := rand.Intn(100) // 生成0到99之间的随机数

	switch {
	case probability < 5:
		return 0
	case probability < 20:
		return 1
	default:
		return 2
	}
}

func GenerateTrainTypeName() string {
	// 设置随机数种子
	rand.Seed(time.Now().UnixNano())

	// 定义可能的火车类型名称
	trainTypes := []string{"GaoTieOne", "GaoTieTwo", "GaoTieSeven", "DongCheOne", "DongCheTen"}

	// 随机选择一个火车类型名称
	MockedTrainTypeName := trainTypes[rand.Intn(len(trainTypes))]

	return MockedTrainTypeName
}

// generateRandomCityName generates a random city name from a predefined list of city names.
func generateRandomCityName() string {
	// Predefined list of city names
	cityNames := []string{
		"nanjing", "shijiazhuang", "wuxi", "shanghaihongqiao", "jiaxingnan",
		"hangzhou", "shanghai", "zhenjiang", "suzhou", "taiyuan",
		"xuzhou", "jinan", "beijing",
	}

	// Seed the random number generator to ensure different results each run
	rand.Seed(time.Now().UnixNano())

	// Generate a random index to pick a city name
	randomIndex := rand.Intn(len(cityNames))

	// Return the randomly selected city name
	return cityNames[randomIndex]
}

func ListToString(stations []string) string {

	// Use a builder for efficient string concatenation
	var builder strings.Builder

	for i, station := range stations {
		if i > 0 {
			builder.WriteString(", ")
		}
		builder.WriteString(fmt.Sprintf("Stations[%d] %s", i, station))
	}

	result := builder.String()
	return result
}

func IntListToString(numbers []int) string {
	// 使用 strings.Builder 进行高效的字符串拼接
	var builder strings.Builder

	for i, number := range numbers {
		if i > 0 {
			builder.WriteString(", ")
		}
		builder.WriteString(fmt.Sprintf("Numbers[%d] %d", i, number))
	}

	result := builder.String()
	return result
}

func StringToList(input string) []string {
	// Split the input string by commas and trim any leading/trailing spaces from each element
	parts := strings.Split(input, ",")
	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}
	return parts
}

// generateRandomTime generates a random time in the format "HH:MM:SS".
func generateRandomTime() string {
	//// Seed the random number generator to ensure different results each run
	//rand.Seed(time.Now().UnixNano())

	// Generate random hours, minutes, and seconds
	hour := rand.Intn(24)   // 0-23
	minute := rand.Intn(60) // 0-59
	second := rand.Intn(60) // 0-59

	// Format the time as "HH:MM:SS"
	return fmt.Sprintf("%02d:%02d:%02d", hour, minute, second)
}

func getRandomTime() string {
	MockedRandomDate := faker.Date()

	////// Seed the random number generator to ensure different results each run
	////rand.Seed(time.Now().UnixNano())
	//// Generate random hours, minutes, and seconds
	//hour := rand.Intn(24)   // 0-23
	//minute := rand.Intn(60) // 0-59
	//second := rand.Intn(60) // 0-59
	//MockedRandomTime := fmt.Sprintf("%02d:%02d:%02d", hour, minute, second)
	MockedRandomTime := generateRandomTime()

	DateAndTime := MockedRandomDate + " " + MockedRandomTime

	return DateAndTime
}

// ConvertCommaSeparatedToBracketed converts a comma-separated string to a bracketed, space-separated string
func ConvertCommaSeparatedToBracketed(input string) string {
	// 删除字符串前后的空白
	input = strings.TrimSpace(input)
	// 按逗号分隔字符串，并去除每个元素前后的空白
	parts := strings.Split(input, ",")
	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}
	// 将分隔后的元素用空格连接，并用方括号包围
	result := fmt.Sprintf("[%s]", strings.Join(parts, " "))
	return result
}

// IntSliceToString converts a slice of integers to a bracketed, space-separated string
func IntSliceToString(ints []int) string {
	// 使用 strings.Builder 进行高效的字符串拼接
	var builder strings.Builder
	builder.WriteString("[")
	for i, val := range ints {
		if i > 0 {
			builder.WriteString(" ")
		}
		builder.WriteString(fmt.Sprintf("%d", val))
	}
	builder.WriteString("]")
	return builder.String()
}

// StringSliceToString converts a slice of strings to a bracketed, space-separated string
func StringSliceToString(strs []string) string {
	// 使用 strings.Builder 进行高效的字符串拼接
	var builder strings.Builder
	builder.WriteString("[")
	for i, val := range strs {
		if i > 0 {
			builder.WriteString(" ")
		}
		builder.WriteString(val)
	}
	builder.WriteString("]")
	return builder.String()
}

// RandomSelectString selects a random string from a given slice of strings
func RandomSelectString(options []string) string {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator
	randomIndex := rand.Intn(len(options))
	return options[randomIndex]
}
