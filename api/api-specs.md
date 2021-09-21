# API Specs

## Product

- [ ] GET /products

- [ ] GET /product/{id}

## Customer

### Admin

- [ ] GET admin/customers

- [ ] GET admin/customer/{id}

- [ ] GET admin/customer/{id}/personalInfo

- [ ] GET admin/customer/{id}/medicalHistory -> admin/medicalHistory?customerId=10

- [ ] GET admin/customer/{id}/vaccination

- [ ] GET admin/customer/{id}/order -> admin/booking?customerId=10

- [ ] PUT admin/customer/{id}/personalInfo

- [ ] PUT admin/customer/{id}/medicalHistory -> admin/medicalHistory?customerId=10

- [ ] PUT admin/customer/{id}/vaccination

- [ ] PUT admin/customer/{id}/order -> admin/booking?customerId=10

- [ ] POST admin/login

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