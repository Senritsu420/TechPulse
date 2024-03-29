import { CONST } from "@/const"

export async function SignUp(email: string, password: string) {
  const res = await fetch(`http://localhost:8080${CONST.SIGN_UP}`, {
    method: "POST",
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({
      email: email,
      password: password
    }),
  })
  const data = await res.json()
  return { status: res.status, data: data }
}

export async function SignIn(email: string, password: string) {
  const res = await fetch(`http://localhost:8080${CONST.SIGN_IN}`, {
    method: "POST",
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({
      email: email,
      password: password
    }),
  })
  const data = await res.json()
  return { status: res.status, data: data }
}

export async function SignOut(email: string, password: string) {
  const res = await fetch(`${CONST.API_BASE_PATH}${CONST.SIGN_OUT}`, {
    method: "POST",
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({
      email: email,
      password: password
    }),
  })
  return { status: res.status }
}
