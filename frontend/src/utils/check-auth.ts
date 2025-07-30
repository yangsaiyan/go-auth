import { fetchData } from "./fetch";

export const checkAuth = async () => {
  return new Promise(async (resolve, reject) => {
    try {
      const response = await fetchData("http://localhost:8080/api/user/check");
      console.log(response);
      resolve(true);
      if (window.location.pathname === "/auth") {
        window.location.href = "/";
      }
    } catch (error) {
      reject(false);
    //   if (window.location.pathname !== "/auth") {
    //     window.location.href = "/auth";
    //   }
    }
  });
};
