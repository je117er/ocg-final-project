# API Specs
## Product
* GET /products
* GET /product/{id}
    
    
## Customer
### Admin
* GET admin/customers
* GET admin/customer/{id}
&nbsp;
* GET admin/customer/{id}/personalInfo
* GET admin/customer/{id}/medicalHistory
* GET admin/customer/{id}/vaccination
* GET admin/customer/{id}/order
&nbsp;
* PUT admin/customer/{id}/personalInfo
* PUT admin/customer/{id}/medicalHistory
* PUT admin/customer/{id}/vaccination
* PUT admin/customer/{id}/order
&nbsp;
* POST admin/login

### Customer
* GET customer/{id}
* PUT customer/{id}
* GET constraints
* POST /checkout
Submit personal info and medical history and vaccination schedule after successfully paying
* POST customer/{id}
* GET customer/{id}/cert

## Clinic
### Admin
* GET admin/clinics
* GET admin/clinic/{id}/customers
* GET admin/clinic/{id}/info
* GET admin/clinic/{id}/schedule
* GET admin/clinic/{id}/inventory
&nbsp;
* PUT admin/clinic/{id}/customers
* PUT admin/clinic/{id}/info
* PUT admin/clinic/{id}/schedule
* PUT admin/clinic/{id}/inventory