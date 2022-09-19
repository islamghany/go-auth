import axios from "../../api/api";
import React, { forwardRef, useRef } from "react";
import styled from "styled-components";
import { useApi, Status } from "../../hooks/useApi";
import { useUserSelector } from "../../context/user";
const StyledInput = styled.input`
  width: 100%;
  border: 1px solid #ccc;
  border-radius: 6px;
  padding: 10px 16px;
  line-height: 1.6;
  margin: 16px;
`;

const StyledForm = styled.div`
  width: 100%;
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
`;
const TextInput = forwardRef(({ name, placeholder, type }, ref) => {
  return (
    <StyledInput ref={ref} name={name} type={type} placeholder={placeholder} />
  );
});

export default function Login() {
  const setUser = useUserSelector((ctx) => ctx.setUser);
  const { status, error, exec } = useApi((e) => {
    console.log(Date.now());
    return axios.post("/token/authenticate/stateless-with-refresh-token", e);
  });
  const emailRef = useRef(null);
  const passwordRef = useRef(null);
  const onSubmit = async (e) => {
    e.preventDefault();
    if (status !== Status.PENDING) {
      const res = await exec({
        email: emailRef.current.value,
        password: passwordRef.current.value,
      });
      if (res.error === null) {
        setUser(res.data.data.user);
      }
    }
  };
  return (
    <StyledForm>
      <form onSubmit={onSubmit}>
        <TextInput
          ref={emailRef}
          name="email"
          placeholder="Enter your email"
          type="email"
        />
        <TextInput
          name="password"
          placeholder="Enter your password"
          type="password"
          ref={passwordRef}
        />
        {status === Status.PENDING ? "Loading..." : <button>submit</button>}
        {error && <p>{error}</p>}
      </form>
    </StyledForm>
  );
}
