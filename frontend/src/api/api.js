import axios from "axios";
import {
  isRefreshTokenEXpired,
  storeRefreshTokenExpiration,
} from "../helpers/refreshToken";

const instant = axios.create({
  baseURL: "http://localhost:4000",
  withCredentials: true,
});

instant.interceptors.response.use(
  (res) => {
    return res;
  },
  async (err) => {
    const originalConfig = err.config;

    console.log("OE : ", originalConfig);
    if (
      originalConfig.url !==
        "/token/authenticate/stateless-with-refresh-token" &&
      err.response
    ) {
      // if the refresh token does't yet expired and the res is 401
      if (
        isRefreshTokenEXpired() &&
        err.response.status === 401 &&
        !originalConfig._retry
      ) {
        originalConfig._retry = true;

        try {
          const res = await instant.post(
            "/token/authenticate/renew-access-token"
          );

          console.log("res", res);
          return instant(originalConfig);
        } catch (_error) {
          console.log(_error);
          return Promise.reject(_error);
        }
      }
    }
    return Promise.reject(err);
  }
);
export default instant;
