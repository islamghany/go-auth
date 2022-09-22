const storeRefreshTokenExpiration = (tokenExpiry) => {
  try {
    localStorage.setItem("tokenExpiry", tokenExpiry);
  } catch (err) {
    console.log(err);
  }
};

const isRefreshTokenEXpired = () => {
  try {
    const tokenExpiry = localStorage.getItem("tokenExpiry");
    console.log(
      tokenExpiry && new Date(tokenExpiry).getTime() > new Date().getTime()
    );
    if (tokenExpiry && new Date(tokenExpiry).getTime() > new Date().getTime()) {
      return true;
    }
  } catch (err) {
    return false;
  }
  return false;
};

export { isRefreshTokenEXpired, storeRefreshTokenExpiration };
