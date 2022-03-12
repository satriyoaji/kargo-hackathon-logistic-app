import React from "react";
import Table from "../../component/Table";

class DriverPage extends React.Component{
    consructor(props){
        super(props);
        this.state = {
            error: null,
            isLoaded: false,
            items: []
        }
    }

    componentDidMount(){
        fetch("https://gorest.co.in/public/v2/users")
        .then(res => res.json())
        .then(
            (result) => {
                this.setState({
                    isLoaded: true, 
                    items: result
                });
            }, 
            (error) => {
                this.setState({
                    isLoaded:true,
                    error
                })
            }
        )
    }

    render(){
        return (
            <div class="container-fluid">
                <h1>test</h1>
            </div>
        )
    }




}