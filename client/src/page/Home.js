import React from "react";
import { Container, Row, Col } from "react-bootstrap";
import { useEffect, useState } from "react";
import axios from "axios";
import { Link } from "react-router-dom";

const Home = () => {

    const [apiData, setApiData] = useState(false);

    useEffect(() => {
        const fetchDate = async() => {
      
            try {

                const apiURL = process.env.REACT_APP_API_ROOT;
                const response = await axios.get(apiURL);

                if(response.status === 200) {
                    if(response?.data?.statusText === "OK"){
                        setApiData(response?.data?.blog_records);
                    }
                }
            }catch(error){
                console.log(error.response);
            }
        };

        fetchDate();

        return () => {};
    },[]);

 

  console.log(apiData);


    return (
        <Container>
            <Row>
                <Col xs="12" className="py-2">
                <h1 className="text-center">
                    My Blog
                </h1>
                </Col>
            
                {apiData && 
                apiData.map((record, index) => (
                    <Col key={index} xs="4" className="py-5 box">
                    
                    <div className="title text-center">
                        <Link to={`/blog/${record.id}`}> 
                            {record.title} 
                        </Link>
                    </div>
                    <div>{record.content}</div>
                        
                    </Col>
                ))}
            </Row>
        </Container>
    )
}

export default Home;