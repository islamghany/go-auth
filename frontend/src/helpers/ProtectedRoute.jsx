import React from "react";
import { Navigate } from "react-router-dom";
import { useUserSelector } from "../context/user";

const getUserRole = (user) => {
  if (user?.id) return ["loggedin"];
  return ["loggedout"];
};

export default function ProtectedRoute({ children, fallback, roles = [] }) {
  const user = useUserSelector((ctx) => ctx.user);

  const cRoles = getUserRole(user);
  if (roles.some((role) => cRoles.includes(role))) {
    return children;
  }
  return <Navigate to={fallback} />;
}
