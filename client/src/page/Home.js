import React from "react";
import { Container, Row, Col, Nav,Navbar, Image } from "react-bootstrap";
import { useEffect, useState } from "react";
import axios from "axios";
import { Link } from "react-router-dom";
import Header from "../component/Header";
import Footer from "../component/Footer";
import Naitigation from "../component/Navitigation";


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
            <Header/>

            
            

            {/* Profile Section */}
            <Row className="my-4">
                <Col className="text-center">
                <Image
                    src="https://via.placeholder.com/100" // 替换为实际头像URL
                    roundedCircle
                />
                <h2>JeaMianh</h2>
                <p>东风夜放花千树，更吹落、星如雨。</p>
                </Col>
            </Row>

            

            {/* Main Content */}
            <Row className="my-4">

                {apiData && 
                apiData.map((record, index) => (
                    <Row>
                        <div key={index} xs="4" className="py-5">
                            <p><strong>{record.created_at}</strong></p>
                            <h3>
                                <Link to={`/blog/${record.id}`}> 
                                    {record.title} 
                                </Link>
                            </h3>
                            <p>{record.content}</p>
                            <hr/>
                        </div>
                    </Row>
                ))}
                
            </Row>

            <Footer/>
        </Container>
    )
}

export default Home;