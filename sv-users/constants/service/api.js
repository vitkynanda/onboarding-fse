import callApi from "./config";

export const getListUsers = async (page) => {
  console.log("exec");
  const url = `http://localhost:8001/users`;
  return callApi({
    url,
    method: "GET",
  });
};

export const getUserDetail = async ({ data }) => {
  const url = `http://localhost:8001/users/${data.id}`;
  return callApi({
    url,
    method: "GET",
  });
};

export const createNewUser = async ({ payload }) => {
  const url = `http://localhost:8001/user`;
  return callApi({
    url,
    method: "POST",
    body: JSON.stringify(payload),
  });
};

export const editUserData = async ({ data, payload }) => {
  const url = `http://localhost:8001/user/${data.id}`;
  return callApi({
    url,
    method: "PUT",
    body: JSON.stringify(payload),
  });
};

export const deleteUserData = async ({ data }) => {
  const url = `http://localhost:8001/user/${data.id}`;
  return callApi({
    url,
    method: "DELETE",
  });
};
