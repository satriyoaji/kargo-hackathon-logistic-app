import {Modal, Button, Form} from "react-bootstrap"
import { useState, useEffect } from "react";
import axios from "axios";

const ModalEditTrucks =(props)=>{
  const [errors, setErrors] = useState({});
  const [truck, setTruck] = useState({});
  const [licenseNumber, setLicenseNumber] = useState("");
  const [prodYear, setProdYear] = useState("");
  const urlPutTruck= "/api/transporter/store-truck/"
  const urlGetTruck= "/api/transporter/get-truck/"

  useEffect(()=>{
    axios.get(urlGetTruck)
    .then((data)=>{
      console.log(data);
      setTruck(data);
    })
    .catch((err)=>{
      console.log(err);
    })
    .finally(()=>{
    })
},[])

  const updateErr = (value) => {
    console.log(!!errors.value, value);
    setErrors({
      ...errors,
      [value]: null,
    });
  };

  const findFormErrors = () => {
    const newErrors = {};
    if (!licenseNumber || licenseNumber === "") newErrors.licenseNumber = "cannot be blank!";
    if (!prodYear || prodYear === "") newErrors.prodYear = "cannot be blank!";

    return newErrors;
  }

  const handleUpdateTruck=(e)=>{
    e.preventDefault();
    const newErrors = findFormErrors();
    if (Object.keys(newErrors).length > 0) {
      setErrors(newErrors);
    } else {
      axios.put(urlPutTruck+props.id)
      .then(()=>{      
      })
      .catch((err)=>{
        console.log(err);
      })
      .finally(()=>{
      })
      props.onHide(e);
    }
  }

  return(
    <>
    <Modal
      {...props}
      size="lg"
      aria-labelledby="contained-modal-title-vcenter"
      centered
    >
      <Modal.Header closeButton>
        <Modal.Title id="contained-modal-title-vcenter">
           Edit Unit
        </Modal.Title>
      </Modal.Header>
      <Modal.Body>
      <Form>
          {/* license Number */}
          <Form.Group className="mb-3" controlId="formBasicEmail">
            <Form.Label>License Number</Form.Label>
            <Form.Control 
                type="text" 
                autoComplete="off"
                placeholder="Enter license number"
                value={truck.license_number}
                onChange={(e) => {
                  setLicenseNumber(e.target.value);
                  updateErr("licenseNumber");
                }}
                required
                isInvalid={!!errors.licenseNumber}
                />
            <Form.Control.Feedback type="invalid">
                {errors.licenseNumber}
            </Form.Control.Feedback>
          </Form.Group>
          {/* License Type */}
          <Form.Group className="mb-3">
            <Form.Label>License Type</Form.Label>
            <Form.Select id="licenseType" value={truck.plate_type}>
              <option value="Black">Black</option>
              <option value="Yellow">Yellow</option>
            </Form.Select>
          </Form.Group>
          <Form.Group className="mb-3">
            <Form.Label>Truck Type</Form.Label>
            <Form.Select id="truckType"   value={truck.truck_type}>
              <option value="Black">Tronton</option>
              <option value="Yellow">Container</option>
              <option value="Yellow">CDE</option>
            </Form.Select>
          </Form.Group>
          <Form.Group className="mb-3" controlId="formBasicEmail">
            <Form.Label>Production Year</Form.Label>
            <Form.Control 
                type="text" 
                placeholder="Enter production year"
                autoComplete="off"
                value={truck.production_year}
                onChange={(e) => {
                  setProdYear(e.target.value);
                  updateErr("prodYear");
                }}
                required
                isInvalid={!!errors.prodYear} />
            <Form.Control.Feedback type="invalid">
                {errors.prodYear}
            </Form.Control.Feedback>
          </Form.Group>
          <Form.Group controlId="formFile" className="mb-3">
            <Form.Label>STNK</Form.Label>
            <Form.Control type="file" />
          </Form.Group>
          <Form.Group controlId="formFile" className="mb-3">
            <Form.Label>KIR</Form.Label>
            <Form.Control type="file" />
          </Form.Group>
        </Form>
      </Modal.Body>
      <Modal.Footer>
        <Button onClick={(e)=>handleUpdateTruck(e)}>Close</Button>
      </Modal.Footer>
    </Modal>
    </>
  )
}
export default ModalEditTrucks;