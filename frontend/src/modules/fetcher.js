import axios from 'axios';

export const ENDPOINT = 'http://localhost:1323'
export const fetcher = async (url) => {
  const rawUserData = sessionStorage.getItem('userData');
  const user = rawUserData ? JSON.parse(rawUserData) : null;

  try {
    const response = await axios(`${ENDPOINT}${url}`,
      {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
          Authorization: 'Bearer ' + user.token,
        }
      });
    console.log("method:",response?.config?.method,"url:",response?.config?.url)
    return response.data;
  } catch (error) {
    console.log(error)
    throw new Error(`Request ${url} failed`);
  }
};


export const mutater = async (url, method, data) => {
  const rawUserData = sessionStorage.getItem('userData');
  const user = rawUserData ? JSON.parse(rawUserData) : null;

  try {
    const response = await axios(`${ENDPOINT}${url}`,
      {
        method: method,
        headers: {
          'Content-Type': 'application/json',
          Authorization: 'Bearer ' + user.token,
        },
        data
      });
    console.log("method:",response?.config?.method,"url:",response?.config?.url)
    return response.data;
  } catch (error) {
    console.log(error)
    throw new Error(`Request ${url} failed`);
  }
};
