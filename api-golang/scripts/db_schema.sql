-- Create the travel planning database
CREATE DATABASE toorme;

-- Switch to the TravelPlanner database
USE toorme;

-- Enable pgcrypto extension for password hashing
CREATE EXTENSION IF NOT EXISTS pgcrypto;

-- 1. Users Table
CREATE TABLE Users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    password TEXT NOT NULL,
    full_name VARCHAR(100),
    email VARCHAR(100) UNIQUE,
    phone_number VARCHAR(15),
    role VARCHAR(20) CHECK (role IN ('admin', 'customer', 'driver')) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 2. Vehicles Table
CREATE TABLE Vehicles (
    id SERIAL PRIMARY KEY,
    driver_id INT REFERENCES Users(id) ON DELETE CASCADE,
    vehicle_type VARCHAR(50),
    license_plate VARCHAR(20) UNIQUE,
    capacity INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 3. Suppliers Table
CREATE TABLE Suppliers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    contact_email VARCHAR(100),
    contact_phone VARCHAR(15),
    address TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 4. Locations Table
CREATE TABLE Locations (
    id SERIAL PRIMARY KEY,
    supplier_id INT REFERENCES Suppliers(id) ON DELETE CASCADE,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    address TEXT,
    latitude DECIMAL(10, 8),
    longitude DECIMAL(11, 8),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 5. Travel Plans Table
CREATE TABLE TravelPlans (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    customer_id INT REFERENCES Users(id) ON DELETE CASCADE,
    status VARCHAR(20) CHECK (status IN ('active', 'inactive')) DEFAULT 'active',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 6. Travel Plan Locations Table
CREATE TABLE TravelPlanLocations (
    id SERIAL PRIMARY KEY,
    travel_plan_id INT REFERENCES TravelPlans(id) ON DELETE CASCADE,
    location_id INT REFERENCES Locations(id) ON DELETE CASCADE,
    visit_order INT NOT NULL,
    UNIQUE (travel_plan_id, location_id)
);

-- 7. Trips Table
CREATE TABLE Trips (
    id SERIAL PRIMARY KEY,
    customer_id INT REFERENCES Users(id) ON DELETE CASCADE,
    driver_id INT REFERENCES Users(id) ON DELETE CASCADE,
    vehicle_id INT REFERENCES Vehicles(id) ON DELETE SET NULL,
    travel_plan_id INT REFERENCES TravelPlans(id) ON DELETE SET NULL,
    location_id INT REFERENCES Locations(id) ON DELETE SET NULL,
    departure_date TIMESTAMP NOT NULL,
    return_date TIMESTAMP,
    status VARCHAR(20) CHECK (status IN ('pending', 'completed', 'canceled')) DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 8. Notifications Table
CREATE TABLE Notifications (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES Users(id) ON DELETE CASCADE,
    message TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    read_at TIMESTAMP
);

-- Sample Data Inserts
-- Insert Users with hashed passwords
INSERT INTO Users (username, password, full_name, email, phone_number, role) 
VALUES 
('admin1', crypt('securepassword123', gen_salt('bf')), 'Admin One', 'admin1@example.com', '123-456-7890', 'admin'),
('customer1', crypt('mypassword456', gen_salt('bf')), 'Customer One', 'customer1@example.com', '234-567-8901', 'customer'),
('driver1', crypt('driverpass789', gen_salt('bf')), 'Driver One', 'driver1@example.com', '345-678-9012', 'driver');

-- Insert Vehicles
INSERT INTO Vehicles (driver_id, vehicle_type, license_plate, capacity) 
VALUES (3, 'Van', 'ABC1234', 15);

-- Insert Suppliers
INSERT INTO Suppliers (name, contact_email, contact_phone, address) 
VALUES 
('Supplier A', 'contact@supplierA.com', '123-123-1234', '123 Main Street, City A'),
('Supplier B', 'contact@supplierB.com', '456-456-4567', '456 High Street, City B');

-- Insert Locations
INSERT INTO Locations (supplier_id, name, description, address, latitude, longitude) 
VALUES 
(1, 'Beach Paradise', 'A beautiful beach with clear water.', 'Beach St, Coastal City', 25.77427, -80.19366),
(1, 'Mountain Retreat', 'A serene mountain escape.', 'Mountain Rd, Alpine Town', 39.73915, -104.99025),
(2, 'Urban Park', 'A bustling city park.', 'Park Ave, Metro City', 40.71278, -74.006),
(2, 'Historic Site', 'A landmark full of history.', 'History Ln, Old Town', 51.50722, -0.12750);

-- Insert Travel Plans
INSERT INTO TravelPlans (name, description, customer_id) 
VALUES ('Coastal Adventure', 'A trip to visit the best coastal spots.', 2);

-- Insert Travel Plan Locations
INSERT INTO TravelPlanLocations (travel_plan_id, location_id, visit_order) 
VALUES (1, 1, 1), (1, 2, 2);

-- Insert Trips
INSERT INTO Trips (customer_id, driver_id, vehicle_id, travel_plan_id, departure_date, return_date) 
VALUES (2, 3, 1, 1, '2024-12-20 08:00:00', '2024-12-25 18:00:00');

-- Insert Notifications
INSERT INTO Notifications (user_id, message) 
VALUES (3, 'You have been assigned a new trip scheduled for 2024-12-20.');
