import { fetchData } from "./fetch";

export const checkAuth = async () => {
  return new Promise(async (resolve, reject) => {
    try {
      const response = await fetchData("http://localhost:8080/api/user/check");
      console.log(response);
      alert("checkAuth success");
      resolve(true);
      if (window.location.pathname === "/auth") {
        window.location.href = "/";
      }
    } catch (error) {
      console.log(error);
      alert("checkAuth error");
      reject(false);
      if (window.location.pathname !== "/auth") {
        window.location.href = "/auth";
      }
    }
  });
};
