# API Specs

## Product

- [x] GET /products

- [x] GET /product/{id}

## Customer

### Admin

- [x] GET admin/customers

- [x] GET admin/customer/{id}

- [x] GET admin/customer/{id}/medicalHistory -> admin/medicalHistory?customerId=10

- [ ] GET admin/customer/{id}/vaccination

- [x] PUT admin/customer

- [ ] PUT admin/customer/{id}/vaccination

- [x] POST admin/login

### Customer

- [x] GET /customer?email=sameple@gmail.com

- [x] PUT /customer

- [x] GET constraints

- [ ] POST /checkout |Submit personal info and medical history and vaccination schedule after successfully paying

- [x] POST /customer

- [ ] GET customer/{id}/cert

## Clinic

### Admin

- [x] GET admin/clinics

- [x] GET admin/clinic/{id}/customers -> admin/customers?clinicId=3x23342tfg4

- [x] GET admin/clinic/{id}/info

- [ ] GET admin/clinic/{id}/schedule -> admin/schedule?clinicId=x3f3f&date=2021-24-15&type=19

- [ ] GET admin/clinic/{id}/inventory -> admin/inventory?clinicId=fjeowfehoe

- [x] PUT admin/clinic/{id}/info

- [ ] PUT admin/clinic/{id}/schedule -> admin/schedule?clinicId=x3f3f&date=2021-24-15&type=19

- [ ] PUT admin/clinic/{id}/inventory -> dmin/inventory?clinicId=fjeowfehoe