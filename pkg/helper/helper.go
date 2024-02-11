package helper

import (
	"context"
	"math"
	"mime/multipart"
	"regexp"
	"time"

	"github.com/aarathyaadhiv/met/pkg/utils/response"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

const diffAge = 5

func IsValidEmail(email string) bool {

	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(pattern, email)
	return match
}

func CalculateAge(dob time.Time) int {
	currentDate := time.Now()

	years := currentDate.Year() - dob.Year()

	if currentDate.Month() < dob.Month() || currentDate.Month() == dob.Month() && currentDate.Day() < dob.Day() {
		years--
	}
	return years
}

func MinAndMaxAge(age int) (int, int) {
	return age - diffAge, age + diffAge
}

func Gender(id uint) uint {
	if id == 1 {
		return id + 1
	} else if id == 2 {
		return id - 1
	}
	return id
}

func HaversineDistance(lat1, lon1, lat2, lon2 float64) float64 {
	const radius = 6371 // Earth radius in kilometers

	dLat := (lat2 - lat1) * (math.Pi / 180.0)
	dLon := (lon2 - lon1) * (math.Pi / 180.0)

	lat1Rad := lat1 * (math.Pi / 180.0)
	lat2Rad := lat2 * (math.Pi / 180.0)

	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Sin(dLon/2)*math.Sin(dLon/2)*math.Cos(lat1Rad)*math.Cos(lat2Rad)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	distance := radius * c

	return distance
}

func AddImageToS3(image *multipart.FileHeader) (string, error) {

	sdkConfig, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("ap-south-1"))
	if err != nil {
		return "", err
	}
	s3Client := s3.NewFromConfig(sdkConfig)

	uploader := manager.NewUploader(s3Client)

	file, err := image.Open()
	if err != nil {

		return "", err
	}
	defer file.Close()
	upload, err1 := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String("fashionstore"),
		Key:    aws.String(image.Filename),
		Body:   file,
		ACL:    "public-read",
	})

	if err1 != nil {
		return "", err1
	}

	return upload.Location, nil
}

func SearchForInterest(a []string, num string, beg int, end int) bool {
	for beg <= end {
		mid := (end + beg) / 2
		if a[mid] == num {
			return true
		} else if a[mid] < num {
			beg = mid + 1
		} else {
			end = mid - 1
		}
	}
	return false
}

func Abs(i int) int {
	if i > 0 {
		return i
	}
	return -i
}

func QuickSort(a []float64, b []response.Home, start, end int) {
	if start < end {
		p := partition(a, b, start, end)
		QuickSort(a, b, start, p-1)
		QuickSort(a, b, p+1, end)
	}
}
func partition(a []float64, b []response.Home, start, end int) int {
	pivot := a[end]
	i := start - 1
	for j := start; j < end; j++ {
		if a[j] < pivot {
			i++
			a[i], a[j] = a[j], a[i]
			b[i], b[j] = b[j], b[i]
		}
	}
	a[i+1], a[end] = a[end], a[i+1]
	b[i+1], b[end] = b[end], b[i+1]
	return i + 1
}
