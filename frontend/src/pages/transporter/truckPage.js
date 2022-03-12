import {Table,Button, Container, DropdownButton, Dropdown, Spinner} from "react-bootstrap"
import { useState, useEffect } from "react"
import axios from "axios"
import ModalEditTrucks from "./editTrucks"
import ModalCreateTrucks from "./addTrucks"
import "./styleTrucks.css"

const TruckPage =()=>{
  const [modalEditShow, setModalEditShow]= useState(false)
  const [modalCreateShow, setModalCreateShow]= useState(false)
  const [loading, setLoading]=useState(true)
  const [trucks, setTrucks]= useState([])
  const [id, setId]= useState([])
  const urlGetTrucks = "/api/tranporter/get-all-trucks"
  const urlDeleteTruck = "/api/transporter/delete-truck/"

  useEffect(()=>{
      setLoading(true)
      axios.get(urlGetTrucks)
      .then((data)=>{
        console.log(data);
        setTrucks(data);
      })
      .catch((err)=>{
        console.log(err);
      })
      .finally(()=>{
        setLoading(false)
      })
  },[])

  const handleDeleteTruck=(idDelete)=>{
    setLoading(true)
    axios.delete(urlDeleteTruck+idDelete)
    .then(()=>{      
    })
    .catch((err)=>{
      console.log(err);
    })
    .finally(()=>{
      setLoading(false)
    })
  }

  if (loading){
    return (
      <>
        <div className="loading">
          <h4>
            Please wait ... <Spinner  animation="grow" />
          </h4>          
        </div>
      </>
    )
  }
  
  return (
    <>
    <Container className="page">
      <Button variant="primary" onClick={()=> setModalCreateShow(true)}>Add New Unit</Button>
      <Table className="table-trucks">
        <thead>
          <tr>
            <th>License Number</th>
            <th>Truck Type</th>
            <th>Plate Type</th>
            <th>Production Year</th>
            <th>Status</th>
            <th>Actions</th>
          </tr>
        </thead>
        <tbody>
            <tr>
              <td>B</td>            
              <td>Trunton</td> 
              <td>Yellow</td> 
              <td>2001</td> 
              <td></td>
              <td>
                  <DropdownButton id="dropdown-basic-button" title="update">
                    <Dropdown.Item onClick={()=>setModalEditShow(true)}>Edit</Dropdown.Item>
                    <Dropdown.Item href="#/action-3">Deactivate Unit</Dropdown.Item>
                    <Dropdown.Item>Delete</Dropdown.Item>
                  </DropdownButton>
                </td>
              </tr>
            {trucks.map((el, idx)=>(
              <tr>
                <td>{el.license_number}</td>            
                <td>{el.truck_type}</td> 
                <td>{el.plate_type}</td> 
                <td>{el.production_year}</td> 
                <td>{el.status}</td>
                <td>
                  <DropdownButton id="dropdown-basic-button" title="update">
                    <Dropdown.Item 
                      onClick={()=>{ 
                        setModalEditShow(true);
                        setId(el.id);
                      }
                     }>Edit</Dropdown.Item>
                    <Dropdown.Item href="#/action-3">Deactivate Unit</Dropdown.Item>
                    <Dropdown.Item onClick={()=>handleDeleteTruck(el.id)}>Delete</Dropdown.Item>
                  </DropdownButton>
                </td>
              </tr>
            ))}
          
        </tbody>
      </Table>
      
      <ModalEditTrucks show={modalEditShow} onHide={()=>setModalEditShow(false)} id={id}/>
      <ModalCreateTrucks show={modalCreateShow} onHide={()=>setModalCreateShow(false)}/>
    </Container>    
    </>
  )

}
export default TruckPage;