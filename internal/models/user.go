package models

type CreateUserRequest struct {
    Name string `json:"name" validate:"required"`
    Dob string `json:"dob"  validate:"required"`
}

type UpdateUserRequest struct {
    Name string `json:"name" validate:"required"`
    Dob string `json:"dob"  validate:"required"`
}

type UserResponse struct {
    ID int32  `json:"id"`
    Name string `json:"name"`
    Dob string `json:"dob"`
}

type UserWithAgeResponse struct {
    ID int32  `json:"id"`
    Name string `json:"name"`
    Dob string `json:"dob"`
    Age int    `json:"age"`
}

type PaginatedUsersResponse struct {
    Users []UserWithAgeResponse `json:"users"`
    Page int `json:"page"`
    Limit int `json:"limit"`
    Total int64 `json:"total"`
}

