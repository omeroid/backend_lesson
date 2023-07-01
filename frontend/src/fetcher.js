import axios from 'axios';

export const fetcher = async (url,options) => {
  try{
    const response = await axios(url,options);
    return response.data;
  }catch(error){
    throw new Error("Request failed");
    console.log(error)
  }
};