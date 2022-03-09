const formatTime = () => {
  const date = new Date();
  const hours = date.getHours();
  const minutes = date.getMinutes();

  return `${hours >= 10 ? hours : "0" + hours} : ${
    minutes >= 10 ? minutes : "0" + minutes
  }`;
};

export default formatTime;
