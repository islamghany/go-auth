import axios from "axios";

const instant = axios.create({
  baseURL: "http://localhost:4000",
  withCredentials: true,
});

export default instant;
