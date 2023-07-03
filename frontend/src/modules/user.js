import { useEffect, useState } from 'react';

export const useUser = () => {
  const [user, setUser] = useState(null);

  useEffect(() => {
    const storageData = sessionStorage.getItem('userData');
    if (storageData) {
      setUser(JSON.parse(storageData))
    }
  }, [])

  return {
    user,
    setUser,
  }

};