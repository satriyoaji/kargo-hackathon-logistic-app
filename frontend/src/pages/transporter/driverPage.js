import React, {Component} from "react";
import 'bootstrap/dist/css/bootstrap.min.css';
import './driver.css';
import AddDriverModal from "./addDriver";
import { Link } from "react-router-dom";

class DriverPage extends React.Component{
    constructor(props){
        super(props);
        this.state = {
            usingSearch: false,
            list_items:3,
            keySearch:null,
            showModalAdd:false, 
            showModalUpdate:false,
            newDriver:null,
            items: [
                {
                    driverName: "Jack",
                    phoneNumber: "+7263726",
                    createdAt: "2 August 2020",
                    status: "Active"
                },
                {
                    driverName: "Hello",
                    phoneNumber: "+7263726",
                    createdAt: "2 August 2020",
                    status: "Active"
                }, 
                {
                    driverName: "Shes",
                    phoneNumber: "+7263726",
                    createdAt: "2 August 2020",
                    status: "Active"
                }
            ],
            filterResult: []
        }
    }
    
    // changeShowModalAdd = () => {
    //     this.state.showModalAdd = !this.state.showModalAdd;
    // }

    // changeShowModalUpdate = () => {
    //     this.state.showModalUpdate = !this.state.showModalUpdate;
    // }

    // getNewDriver = (driver) => {
    //     this.state.showModalAdd = !(this.state.showModalAdd);
    //     this.state.newDriver = driver;
    // }

    handleFilter = (val) => {
        if(val != ''){
            this.setState({keySearch: val, usingSearch:true});
            const filteredData = this.state.items.filter(item => {
                return item.driverName.toLowerCase().startsWith(val.toLowerCase());
            })
            this.setState({filterResult: filteredData});
        }
        else{
            this.setState({keySearch: null, usingSearch:false})
        }
    }
    
    render(){
        const {items, list_items, keySearch, usingSearch, filterResult, showModalAdd} = this.state;
        return (
            <div class="container">
            <div class="container-fluid box" style={{marginTop: "50px"}}>
                <div class="d-flex justify-content-between top-content" str>
                    <div class="col-md-6 mb-4">
                        <div class="row">
                            <div class="col-5">
                                <input value={keySearch} onChange={(e)=> this.handleFilter(e.target.value)} placeholder="Search..."></input>
                            </div>
                            <div class="col-5">
                                <Link to="/login/driver/add" class="button">Add Driver</Link>
                            </div>
                        </div>
                    </div>
                </div>
                {list_items == null && 
                    <div>
                        <span class="sr-only">Loading</span>
                    </div>
                }
                {list_items != null &&
                    <div class="tab-content">
                        <table class="table table-borderless">
                            <thead>
                                <tr>
                                    <th scope="col">Driver Name</th>
                                    <th scope="col">Phone Number</th>
                                    <th scope="col">Created At</th>
                                    <th scope="col">Status</th>
                                    <th scope="col"></th>
                                </tr>
                            </thead>
                            {!usingSearch && 
                                <tbody>
                                    {items.map(item => (
                                        <tr>
                                            <td scope="row">{item.driverName}</td>
                                            <td>{item.phoneNumber}</td>
                                            <td>{item.createdAt}</td>
                                            <td>{item.status}</td>
                                            <td><button>Update</button></td>
                                        </tr>
                                    ))}
                                </tbody>
                            }
                            {usingSearch  && 
                                <tbody>
                                    {filterResult.map(item => (
                                        <tr>
                                            <td scope="row">{item.driverName}</td>
                                            <td>{item.phoneNumber}</td>
                                            <td>{item.createdAt}</td>
                                            <td>{item.status}</td>
                                            <td><button>Update</button></td>
                                        </tr>
                                    ))}
                                </tbody>
                            }
                        </table>
                    </div>
                }
            </div>
            </div>
        )
    }
}

export default DriverPage;