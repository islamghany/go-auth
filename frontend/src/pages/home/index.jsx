import React from "react";
import { Link } from "react-router-dom";
import styled from "styled-components";

const HomeContainer = styled.div`
  width: 100%;
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  .links {
    display: flex;
    a {
      margin: 0 4px;
    }
  }
`;
export default function Home() {
  return (
    <HomeContainer>
      <div className="links">
        <Link to="/login">Login</Link>
        <Link to="/protected-page">ProtectedPage</Link>
      </div>
    </HomeContainer>
  );
}
