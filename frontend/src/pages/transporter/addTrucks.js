import {Modal, Button, Form} from "react-bootstrap"
import { useState } from "react";
import axios from "axios";

const ModalCreateTrucks =(props)=>{
  const [form, setForm] = useState({});
  const [errors, setErrors] = useState({});
  const [plateType, setPlateType] = useState ("")
  const [truckType, setTruckType] = useState ("")
  const { licenseNumber, prodYear } = form;
  const urlPostTruck= "/api/transporter/store-truck";
  const setField = (field, value) => {
    setForm({
      ...form,
      [field]: value,
    });
    // Check and see if errors exist, and remove them from the error object:
    if (!!errors[field])
      setErrors({
        ...errors,
        [field]: null,
      });
  };

  const findFormErrors = () => {
    const newErrors = {};
    if (!licenseNumber || licenseNumber === "") newErrors.licenseNumber = "cannot be blank!";
    if (!prodYear || prodYear === "") newErrors.prodYear = "cannot be blank!";

    return newErrors;
  }

  const handleCreateTruck=(e)=>{
    e.preventDefault();
    const newErrors = findFormErrors();
    if (Object.keys(newErrors).length > 0) {
      setErrors(newErrors);
    } else {
      const body = {
        license_number : licenseNumber,
        truck_type : truckType,
        plate_type : plateType,
        production_year : prodYear
      }
      console.log(body);
      axios.delete(urlPostTruck, body)
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
          Add New Unit
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
                onChange={(e) => setField("licenseNumber", e.target.value)}
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
            <Form.Select id="licenseType" onChange={(e) => setPlateType(e.target.value)}>
              <option value="Black">Black</option>
              <option value="Yellow">Yellow</option>
            </Form.Select>
          </Form.Group>
          <Form.Group className="mb-3">
            <Form.Label>Truck Type</Form.Label>
            <Form.Select id="truckType" onChange={(e) => setTruckType(e.target.value)}>
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
                onChange={(e) => setField("prodYear", e.target.value)}
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
        <Button onClick={(e)=>handleCreateTruck(e)}>Close</Button>
      </Modal.Footer>
    </Modal>
    </>
  )
}
export default ModalCreateTrucks;