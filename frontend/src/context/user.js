import { useCallback, useMemo, useState } from "react";
import contextFactory from "./contextFactory";

const [UContext, useUserContext, useUserSelector] = contextFactory();

const UserContext = ({ children }) => {
  const [user, setUser] = useState({ id: -1 });

  let u = useMemo(() => user, []);

  const setUserAction = useCallback((e) => setUser(e), [setUser]);

  return (
    <UContext.Provider value={{ user, setUser: setUserAction }}>
      {children}
    </UContext.Provider>
  );
};

export { useUserSelector };

export default UserContext;
