package messages

type Application struct {
	ID   int
	Name string
}

type User struct {
	ID           int
	EmailAddress string
	FirstName    string
	LastName     string
}

type ListApplicationsResponse struct {
	Result []Application
}

type Login struct {
	EmailAddress       string
	Password           string
	OrganizationUnitID int
	ApplicationID      int
}

type LoginResponse struct {
	Authentication         int
	OrganizationUnits      []interface{}
	User                   interface{}
	LoggedInOrganizationID int

	AuthenticationToken string
}

type StoreBlob struct {
	Category     string
	OriginalName string
	MimeType     string
	Data         []byte
}

type StoreBlobResponse struct {
	Guid string
	Url  string
}

type GetBlobInfoResponse struct {
	Guid         string
	MimeType     string
	OriginalName string
	Category     string
	Size         int
	Url          string
}

type GetCurrentUserResponse struct {
	User User
}

type GetProductDetailResponse struct {
	Result map[string]interface{}
}

type GetProductByBarcode struct {
	Barcode string
}
