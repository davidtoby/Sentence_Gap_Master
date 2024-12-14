import axios from 'axios'

const urlPrefix = 'http://localhost'

export const choose = (id) => {
    return axios.get(`${urlPrefix}/choose`, {
        params: {
            id,
        },
    })
}

export const submit = (msg) => {
    return axios.post(`${urlPrefix}/submit`, msg)
}
