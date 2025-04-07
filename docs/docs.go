// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": ["http"],
    "swagger": "2.0",
    "info": {
        "description": "This is a sample API for managing Admin users.",
        "title": "Admin API",
        "contact": {},
        "version": "1.0"
    },
    "host": "localhost:8080",
    "basePath": "/",
    "paths": {
        "/admin/signup": {
            "post": {
                "tags": ["admin"],
                "summary": "Admin Signup",
                "description": "Admin user signs up for the system",
                "parameters": [
                    {
                        "name": "admin",
                        "in": "body",
                        "description": "Admin Signup Details",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/Admin"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Admin created successfully",
                        "schema": {
                            "$ref": "#/definitions/AdminResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request body"
                    },
                    "500": {
                        "description": "Internal server error"
                    }
                }
            }
        },
        "/admin/login": {
            "post": {
                "tags": ["admin"],
                "summary": "Admin Login",
                "description": "Admin user logs in to the system",
                "parameters": [
                    {
                        "name": "admin",
                        "in": "body",
                        "description": "Admin Login Details",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/Admin"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Admin logged in successfully",
                        "schema": {
                            "$ref": "#/definitions/AdminResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request body"
                    },
                    "500": {
                        "description": "Internal server error"
                    }
                }
            }
        },
        "/admin/logout": {
            "post": {
                "tags": ["admin"],
                "summary": "Admin Logout",
                "description": "Admin user logs out from the system",
                "parameters": [
                    {
                        "name": "admin",
                        "in": "body",
                        "description": "Admin logout Details",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/Admin"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Admin logged out successfully",
                        "schema": {
                            "$ref": "#/definitions/AdminResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request body"
                    },
                    "500": {
                        "description": "Internal server error"
                    }
                }
            }
        }
    },
     "/bookings": {
            "post": {
                "tags": ["Bookings"],
                "summary": "Create a new booking",
                "description": "Create a new booking for a room",
                "parameters": [
                    {
                        "name": "booking",
                        "in": "body",
                        "description": "Booking details to create a new booking",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/Booking"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Booking created successfully",
                        "schema": {
                            "$ref": "#/definitions/BookingResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid booking data"
                    },
                    "404": {
                        "description": "Room not found"
                    },
                    "409": {
                        "description": "Room already occupied"
                    },
                    "500": {
                        "description": "Internal server error"
                    }
                }
            }
        },
        "/bookings": {
            "get": {
                "tags": ["Bookings"],
                "summary": "Get all bookings",
                "description": "Retrieve all bookings with optional sorting",
                "parameters": [
                    {
                        "name": "sortbytime",
                        "in": "query",
                        "description": "Sort bookings by time (asc/desc)",
                        "required": false,
                        "type": "string"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Bookings retrieved successfully",
                        "schema": {
                            "$ref": "#/definitions/BookingListResponse"
                        }
                    },
                    "400": {
                        "description": "Failed to retrieve bookings"
                    },
                    "500": {
                        "description": "Internal server error"
                    }
                }
            }
        },
        "/bookings/{id}/status": {
            "put": {
                "tags": ["Bookings"],
                "summary": "Update booking status",
                "description": "Update the status of a booking (approved/rejected)",
                "parameters": [
                    {
                        "name": "id",
                        "in": "path",
                        "description": "Booking ID",
                        "required": true,
                        "type": "string"
                    },
                    {
                        "name": "status",
                        "in": "query",
                        "description": "Booking status (approved/rejected)",
                        "required": true,
                        "type": "string"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Booking status updated successfully",
                        "schema": {
                            "$ref": "#/definitions/BookingStatusResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid status value"
                    },
                    "404": {
                        "description": "Booking not found"
                    },
                    "500": {
                        "description": "Internal server error"
                    }
                }
            }
        }
    },
     "/approve": {
            "get": {
                "tags": ["Bookings"],
                "summary": "Approve a booking",
                "description": "Approve a booking by updating the status to 'Confirmed'",
                "parameters": [
                    {
                        "name": "managerID",
                        "in": "query",
                        "description": "Manager's ID for validation",
                        "required": true,
                        "type": "string"
                    },
                    {
                        "name": "bookingID",
                        "in": "query",
                        "description": "Booking ID to approve",
                        "required": true,
                        "type": "string"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Booking approved successfully",
                        "schema": {
                            "$ref": "#/definitions/BookingStatusResponse"
                        }
                    },
                    "400": {
                        "description": "Missing bookingID or managerID"
                    },
                    "500": {
                        "description": "Internal server error"
                    }
                }
            }
        },
        "/reject": {
            "get": {
                "tags": ["Bookings"],
                "summary": "Reject a booking",
                "description": "Reject a booking by updating the status to 'Rejected'",
                "parameters": [
                    {
                        "name": "managerID",
                        "in": "query",
                        "description": "Manager's ID for validation",
                        "required": true,
                        "type": "string"
                    },
                    {
                        "name": "bookingID",
                        "in": "query",
                        "description": "Booking ID to reject",
                        "required": true,
                        "type": "string"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Booking rejected successfully",
                        "schema": {
                            "$ref": "#/definitions/BookingStatusResponse"
                        }
                    },
                    "400": {
                        "description": "Missing bookingID or managerID"
                    },
                    "500": {
                        "description": "Internal server error"
                    }
                }
            }
        }
    },
    "definitions": {
        "Admin": {
            "type": "object",
            "properties": {
                "first_name": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                }
            }
        },
        "AdminResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "data": {
                    "type": "object"
                }
            }
        }
    },
    "Booking": {
            "type": "object",
            "properties": {
                "booking_id": {
                    "type": "string"
                },
                "guest_id": {
                    "type": "string"
                },
                "room_number": {
                    "type": "string"
                },
                "number_of_guests": {
                    "type": "integer",
                    "example": 2
                },
                "food_items": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "status": {
                    "type": "string"
                },
                "created_time": {
                    "type": "string",
                    "format": "date-time"
                },
                "updated_time": {
                    "type": "string",
                    "format": "date-time"
                }
            }
        },
        "BookingResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "data": {
                    "$ref": "#/definitions/Booking"
                }
            }
        },
        "BookingListResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/Booking"
                    }
                }
            }
        },
        "BookingStatusResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "data": {
                    "type": "string",
                    "example": "Booking status updated successfully"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
