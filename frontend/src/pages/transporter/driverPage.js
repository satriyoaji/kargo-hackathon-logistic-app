import React, {Component} from "react";
import 'bootstrap/dist/css/bootstrap.min.css';
import './driver.css';

class DriverPage extends React.Component{
    constructor(props){
        super(props);
        this.state = {
            usingSearch: false,
            list_items:3,
            keySearch:'Search keywords...',
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
        }
    }

    handleFilter = (event) => {
        this.setState({keySearch: event.target.value});
    }
    
    render(){
        const {items, list_items, keySearch} = this.state;
        const lowerCasedFilter = keySearch.toLowerCase();
        const filteredData = items.filter(item => {
            return item.driverName.toLowerCase().startsWith(lowerCasedFilter)
        })
        return (
            <div class="container-fluid" style={{marginTop: "50px"}}>
                <div class="d-flex justify-content-between">
                    <div class="col-md-6 mb-4">
                        <div class="row">
                            <div class="col-5">
                                <input value={keySearch} onChange={this.handleFilter}></input>
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
                    <div class="container-fluid tab-content">
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
                            
                        </table>
                    </div>
                }
                
            </div>

        )
    }
}

export default DriverPage;