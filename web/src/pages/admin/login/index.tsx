import { useRef, useState } from "react"
import { login } from "../../../api"

const Login = () => {
  const passwordRef = useRef<HTMLInputElement>(null)
  const [borderColor, setBorderColor] = useState("border-zinc-500")
  const [processing, setProcessing] = useState(false)

  const onLoginPress = () => {
    setProcessing(true)
    const onError = (err: any) => {
      setProcessing(false)
      setBorderColor("border-red-600")
      console.warn(err)
    }
    const onSuccess = (ok: boolean) => {
      if (ok) {
        location.reload()
      } else {
        onError("Incorrect password!")
      }
    }
    login(passwordRef.current!.value).then(onSuccess).catch(onError)
  }

  return (
    <div className="flex flex-1 items-center justify-center">
      <div className="bg-zinc-700 p-4 rounded-lg">
        <h1 className="flex flex-1 justify-center">
          Fazer login
        </h1>
        <input
          ref={passwordRef}
          placeholder="Senha..."
          type="password"
          className={"m-2 p-1 rounded-md bg-zinc-600 border-2 " + borderColor}
          onKeyUp={e => e.key == "Enter" && onLoginPress()}
        />
        <button
          className={"p-1.5 rounded-lg " + (processing ? "bg-gray-500" : "bg-blue-600")}
          disabled={processing}
          onClick={onLoginPress}
        >
          Entrar
        </button>
      </div>
    </div>
  )
}

export default Login
