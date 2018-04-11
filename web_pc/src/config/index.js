/**
 * 项目配置信息
 */
const isDev = () => {
    return window.location.host.indexOf("localhost") !== -1;
};

let def = {
    API: "http://localhost:8888/",
};

let dev = {

};

let prod = {

};

let config = isDev() ? dev : prod;
config = Object.assign(config, def);
export { config }
