import axios from "../api/api";
import React from "react";
import { Status, useApi } from "../hooks/useApi";
import { useUserSelector } from "../context/user";

export default function WithAuth({ children }) {
  const setUser = useUserSelector((ctx) => ctx.setUser);
  const { status } = useApi(
    () => {
      console.log(Date.now());
      return axios.get("/get-user").then((res) => res.data);
    },
    {
      enabled: true,
      onSuccess: (data) => {
        console.log(Date.now(), data);
        setUser(data.user);
      },
    }
  );
  if (status === Status.IDLE || status === Status.PENDING) {
    return <h1>Loading.....</h1>;
  }
  return children;
}
