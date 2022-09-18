import { useEffect, useState } from "react";

export const Status = {
  PENDING: "PENDING",
  ERROR: "ERROR",
  SUCCESS: "SUCCESS",
  IDLE: "IDLE",
};

export const useApi = (fn, config = {}) => {
  const { initialData, enabled, onSuccess, onError, onSettled, initialArgs } =
    config;

  const [data, setData] = useState(initialData);
  const [error, setError] = useState(null);
  const [status, setStatus] = useState(Status.IDLE);
  const exec = async (args) => {
    try {
      setStatus(Status.PENDING);
      const data = await fn(args);
      setData(data);
      setStatus(Status.SUCCESS);
      onSuccess?.(data);
      onSettled?.(data, null);

      return {
        data,
        error: null,
      };
    } catch (err) {
      const error = err?.response?.data.error || "NEWOEN ERROR!";
      setError(error);
      setStatus(Status.ERROR);
      onError?.(error);
      onSettled?.(null, error);
      return {
        data: null,
        error,
      };
    }
  };

  useEffect(() => {
    if (enabled) {
      exec(initialArgs);
    }
  }, []);

  return {
    exec,
    data,
    setData,
    error,
    setError,
    status,
    setStatus,
  };
};
