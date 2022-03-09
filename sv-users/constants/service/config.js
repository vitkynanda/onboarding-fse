import { toast } from "react-toastify";
const callApi = async ({ url, body, method }) => {
  try {
    let headers = { "Content-Type": "application/json" };
    const response = await fetch(url, {
      method,
      headers,
      body,
    });

    if (response.status === 204) {
      toast.success("Deleted data successfully");
    }

    return response.json();
  } catch (err) {}
};

export default callApi;
