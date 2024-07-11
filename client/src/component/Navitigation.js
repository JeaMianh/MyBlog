import React from "react";
import { Navbar, Nav } from "react-bootstrap";

const Naitigation = () => {

    return (
       
        <Navbar bg="light" expand="lg">
            <Navbar.Toggle aria-controls="basic-navbar-nav" />
            <Navbar.Collapse id="basic-navbar-nav">
                <Nav className="ml-auto">
                    <Nav.Link href="#home">Home</Nav.Link>
                    <Nav.Link href="#archive">Archive</Nav.Link>
                    <Nav.Link href="#tags">Tags</Nav.Link>
                    <Nav.Link href="#about">About</Nav.Link>
                </Nav>
            </Navbar.Collapse>
        </Navbar>
    );
}

export default Naitigation;