import React from "react";
import { Link } from "react-router-dom";
import { styled } from "styled-components";

// assets
// import { ReactComponent as AboutIcon } from "././public/assets/database.svg";

//Import necessary FontAwesome components
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faInfoCircle } from "@fortawesome/free-solid-svg-icons";

const Container = styled.div`
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  backgound-size: cover;
  background-color: #18a19c;
`;

const Icon = styled.div`
  justify-content: right;
  align-items: right;
`;
const Title = styled.h1`
  font-size: 2rem;
  color: #333;
`;

const Homepage: React.FC = () => {
  return (
    <Container>
      <Title>Homepage</Title>
      <br></br>
      <Icon>
        <img src="../../public/assets/database.svg" alt="" />
      </Icon>
      <br></br>
      <div>
        <Link to="/about">
          <FontAwesomeIcon icon={faInfoCircle} />
        </Link>
        <Link to="/login">
          <FontAwesomeIcon icon={faInfoCircle} />
        </Link>
      </div>
    </Container>
  );
};

export default Homepage;
