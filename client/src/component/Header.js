import React from "react";
import { Container, Row, Col } from "react-bootstrap";
import Navitigation from "./Navitigation";

const Header = () => {

    return (
       
        <Row className="my-4">
            <Navitigation/>
        </Row>
    );
}

export default Header;