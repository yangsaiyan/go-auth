import axios from "axios";

export const fetchData = async (url: string) => {
  return new Promise(async (resolve, reject) => {
    try {
      const response = await axios.get(url, {
        withCredentials: true,
      });
      resolve(response.data);
    } catch (error) {
      reject(error);
    }
  });
};

export const postData = async (url: string, body: any) => {
  return new Promise(async (resolve, reject) => {
    try {
      const response = await axios.post(url, body, {
        withCredentials: true,
      });
      resolve(response.data);
    } catch (error) {
      reject(error);
    }
  });
};
