import React, { Component } from "react";

import logo from "../../assets/logo.svg";
import "./styles.css";

export default class Main extends Component {

    handleSubmit = async e => {
        e.preventDefault();

        // chama api de login?
        const id = 'id'
        console.log(id)
        this.props.history.push(`/upl/${id}`)
    }

    render() {
        return (
            <div id="main-container">
                <form onSubmit={this.handleSubmit} >
                <img src={logo} alt=""/>
                <input
                    placeholder="Usuário"
                />
                <button type="submit">Login</button>
                </form>
            </div>
        );
    }
}
