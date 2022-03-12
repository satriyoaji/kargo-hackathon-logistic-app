import React, { useState } from "react";
import { Link } from "react-router-dom";
// import { Driver } from "./driverModel";
import './driver.css';

const AddDriverModal = ({getShowAddModal, getNewDriver}) =>{
    const [addForm, setAddForm] = useState({
        driverName: '',
        phoneNumber: '',
    })

    const handleChangeForm = (event) => {
        const nameField = event.target.getAttribute('name');
        const valueField = event.target.value;
        const newData = { ...addForm };
        newData[nameField] = valueField;
        setAddForm(newData);
    }

    const submitForm = (event) => {
        const newDriver = {
            shipment: addForm.driverName,
            phoneNumber: addForm.phoneNumber
        }
        //belom dikirim
    }

    return (
        <div class="container-fluid" style={{marginTop:"50px"}}>
            <div class="container-fluid">
                <form>
                    <div class="row" style={{marginBottom: '20px'}}>
                        <div class="col-lg-8 col-sm-12">
                            <input type='text' name="driverName" required='required' placeholder="Enter driver's name..." onChange={handleChangeForm}></input>
                        </div>
                    </div>
                    <div class="row" style={{marginBottom: '20px'}}>
                        <div class="col-lg-8 col-sm-12">
                            <input type='text' name="phoneNumber" required='required' placeholder="Enter phone number" onChange={handleChangeForm}></input>
                        </div>
                    </div>
                    <div class="row">
                        <div class="col-lg-8 col-sm-12">
                            <Link to="/login/driver" class="button">Back</Link>
                        </div>
                        <div class="col-lg-8 col-sm-12">
                            <Link to="/login/driver" class="button">Submit</Link>
                        </div>
                    </div>
                </form>
            </div>
        </div>
    )
}
export default AddDriverModal