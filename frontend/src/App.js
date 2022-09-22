import "./App.css";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import Home from "./pages/home";
import Login from "./pages/login";
import ProtectedRoute from "./helpers/ProtectedRoute";
import UserContext from "./context/user";
import WithAuth from "./helpers/WithAuth";
import { useApi } from "./hooks/useApi";
import axios from "./api/api";
const PP = () => {
  const { status } = useApi(
    () => axios.get("/protected-resources").then((res) => res.data),
    {
      enabled: true,
      onSettled: (e) => console.log(e),
    }
  );
  return <h1>This is Protected Path {status}</h1>;
};
function App() {
  return (
    <UserContext>
      <WithAuth>
        <Router>
          <Routes>
            <Route index element={<Home />} />
            <Route
              path="/login"
              element={
                <ProtectedRoute roles={["loggedout"]} fallback="/">
                  <Login />
                </ProtectedRoute>
              }
            />
            <Route
              path="/protected-page"
              element={
                <ProtectedRoute roles={["loggedin"]} fallback="/">
                  <PP />
                </ProtectedRoute>
              }
            />
          </Routes>
        </Router>
      </WithAuth>
    </UserContext>
  );
}

export default App;
