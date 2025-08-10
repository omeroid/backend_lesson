import * as React from 'react'
import { useState } from 'react'
import { useNavigate, Link } from 'react-router-dom'
import { motion } from 'framer-motion'
import axios from 'axios'
import toast from 'react-hot-toast'
import { ENDPOINT } from '../modules/fetcher'

export default function SignIn() {
  const navigate = useNavigate()
  const [loading, setLoading] = useState(false)
  const [formData, setFormData] = useState({
    username: '',
    password: ''
  })

  const handleSubmit = async (event) => {
    event.preventDefault()
    setLoading(true)
    
    const url = ENDPOINT + '/user/signin'

    try {
      const response = await axios.post(url, {
        userName: formData.username,
        password: formData.password,
      })
      sessionStorage.setItem('userData', JSON.stringify(response.data))
      toast.success('ログイン成功！チャットへ移動します...')
      setTimeout(() => navigate('/chat'), 500)
    } catch (e) {
      toast.error('ログインに失敗しました。認証情報を確認してください。')
      console.error('Login error:', e?.response)
    } finally {
      setLoading(false)
    }
  }

  const handleChange = (e) => {
    setFormData({
      ...formData,
      [e.target.name]: e.target.value
    })
  }

  return (
    <div style={{ 
      minHeight: '100vh', 
      display: 'flex', 
      alignItems: 'center', 
      justifyContent: 'center', 
      padding: '1rem',
      background: 'linear-gradient(135deg, #000000 0%, #1a1a1a 100%)'
    }}>
      <motion.div
        initial={{ opacity: 0, y: 20 }}
        animate={{ opacity: 1, y: 0 }}
        transition={{ duration: 0.5 }}
        style={{ width: '100%', maxWidth: '28rem' }}
      >
        <div style={{ 
          borderRadius: '1.5rem', 
          padding: '2rem', 
          boxShadow: '0 20px 25px -5px rgba(0, 0, 0, 0.5)',
          background: 'rgba(255, 255, 255, 0.03)',
          backdropFilter: 'blur(20px)',
          border: '1px solid rgba(255, 255, 255, 0.1)'
        }}>
          <motion.div
            initial={{ scale: 0.5 }}
            animate={{ scale: 1 }}
            transition={{ 
              delay: 0.2,
              type: "spring",
              stiffness: 260,
              damping: 20
            }}
            style={{ display: 'flex', justifyContent: 'center', marginBottom: '2rem' }}
          >
            <div style={{ position: 'relative' }}>
              {/* omeroid Logo */}
              <img 
                src="/omeroid.svg" 
                alt="omeroid" 
                style={{ 
                  width: '5rem', 
                  height: '5rem',
                  borderRadius: '1rem',
                  boxShadow: '0 10px 20px rgba(0, 0, 0, 0.3)'
                }}
              />
            </div>
          </motion.div>

          <motion.h2
            initial={{ opacity: 0 }}
            animate={{ opacity: 1 }}
            transition={{ delay: 0.3 }}
            style={{ 
              fontSize: '1.875rem', 
              fontWeight: 'bold', 
              textAlign: 'center', 
              marginBottom: '0.5rem',
              color: '#ffffff'
            }}
          >
            omeroid chat
          </motion.h2>
          <motion.p
            initial={{ opacity: 0 }}
            animate={{ opacity: 1 }}
            transition={{ delay: 0.4 }}
            style={{ textAlign: 'center', color: 'rgb(156, 163, 175)', marginBottom: '2rem' }}
          >
            アカウントにログインして続ける
          </motion.p>

          <form onSubmit={handleSubmit} style={{ display: 'flex', flexDirection: 'column', gap: '1.5rem' }}>
            <motion.div
              initial={{ x: -50, opacity: 0 }}
              animate={{ x: 0, opacity: 1 }}
              transition={{ delay: 0.5 }}
            >
              <label style={{ display: 'block', fontSize: '0.875rem', fontWeight: '500', color: 'rgb(209, 213, 219)', marginBottom: '0.5rem' }}>
                ユーザー名
              </label>
              <div style={{ position: 'relative' }}>
                <div style={{ position: 'absolute', top: '50%', transform: 'translateY(-50%)', left: '0.75rem', pointerEvents: 'none' }}>
                  <svg style={{ width: '1.25rem', height: '1.25rem', color: 'rgb(107, 114, 128)' }} fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
                  </svg>
                </div>
                <input
                  type="text"
                  name="username"
                  value={formData.username}
                  onChange={handleChange}
                  required
                  style={{
                    display: 'block',
                    width: '100%',
                    paddingLeft: '2.5rem',
                    paddingRight: '0.75rem',
                    paddingTop: '0.75rem',
                    paddingBottom: '0.75rem',
                    backgroundColor: 'rgba(255, 255, 255, 0.05)',
                    border: '1px solid rgba(255, 255, 255, 0.1)',
                    borderRadius: '0.75rem',
                    color: 'white',
                    outline: 'none',
                    transition: 'all 0.2s'
                  }}
                  placeholder="ユーザー名を入力"
                  onFocus={(e) => {
                    e.target.style.border = '1px solid rgba(255, 255, 255, 0.3)'
                    e.target.style.backgroundColor = 'rgba(255, 255, 255, 0.08)'
                  }}
                  onBlur={(e) => {
                    e.target.style.border = '1px solid rgba(255, 255, 255, 0.1)'
                    e.target.style.backgroundColor = 'rgba(255, 255, 255, 0.05)'
                  }}
                />
              </div>
            </motion.div>

            <motion.div
              initial={{ x: -50, opacity: 0 }}
              animate={{ x: 0, opacity: 1 }}
              transition={{ delay: 0.6 }}
            >
              <label style={{ display: 'block', fontSize: '0.875rem', fontWeight: '500', color: 'rgb(209, 213, 219)', marginBottom: '0.5rem' }}>
                パスワード
              </label>
              <div style={{ position: 'relative' }}>
                <div style={{ position: 'absolute', top: '50%', transform: 'translateY(-50%)', left: '0.75rem', pointerEvents: 'none' }}>
                  <svg style={{ width: '1.25rem', height: '1.25rem', color: 'rgb(107, 114, 128)' }} fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z" />
                  </svg>
                </div>
                <input
                  type="password"
                  name="password"
                  value={formData.password}
                  onChange={handleChange}
                  required
                  style={{
                    display: 'block',
                    width: '100%',
                    paddingLeft: '2.5rem',
                    paddingRight: '0.75rem',
                    paddingTop: '0.75rem',
                    paddingBottom: '0.75rem',
                    backgroundColor: 'rgba(255, 255, 255, 0.05)',
                    border: '1px solid rgba(255, 255, 255, 0.1)',
                    borderRadius: '0.75rem',
                    color: 'white',
                    outline: 'none',
                    transition: 'all 0.2s'
                  }}
                  placeholder="パスワードを入力"
                  onFocus={(e) => {
                    e.target.style.border = '1px solid rgba(255, 255, 255, 0.3)'
                    e.target.style.backgroundColor = 'rgba(255, 255, 255, 0.08)'
                  }}
                  onBlur={(e) => {
                    e.target.style.border = '1px solid rgba(255, 255, 255, 0.1)'
                    e.target.style.backgroundColor = 'rgba(255, 255, 255, 0.05)'
                  }}
                />
              </div>
            </motion.div>

            <motion.div
              initial={{ y: 20, opacity: 0 }}
              animate={{ y: 0, opacity: 1 }}
              transition={{ delay: 0.7 }}
            >
              <button
                type="submit"
                disabled={loading}
                style={{
                  width: '100%',
                  padding: '0.75rem 1rem',
                  background: loading ? 'rgb(75, 85, 99)' : '#ffffff',
                  color: loading ? 'white' : '#000000',
                  fontWeight: '600',
                  borderRadius: '0.75rem',
                  border: 'none',
                  cursor: loading ? 'not-allowed' : 'pointer',
                  opacity: loading ? 0.5 : 1,
                  display: 'flex',
                  alignItems: 'center',
                  justifyContent: 'center',
                  gap: '0.5rem',
                  transition: 'all 0.2s',
                  transform: 'scale(1)',
                }}
                onMouseEnter={(e) => !loading && (e.currentTarget.style.transform = 'scale(1.02)')}
                onMouseLeave={(e) => !loading && (e.currentTarget.style.transform = 'scale(1)')}
              >
                {loading ? (
                  <div style={{ 
                    width: '1.25rem', 
                    height: '1.25rem', 
                    border: '2px solid white', 
                    borderTopColor: 'transparent', 
                    borderRadius: '50%', 
                    animation: 'spin 1s linear infinite' 
                  }} />
                ) : (
                  <>
                    ログイン
                    <svg style={{ width: '1.25rem', height: '1.25rem' }} fill="none" viewBox="0 0 24 24" stroke="currentColor">
                      <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M14 5l7 7m0 0l-7 7m7-7H3" />
                    </svg>
                  </>
                )}
              </button>
            </motion.div>
          </form>

          <motion.div
            initial={{ opacity: 0 }}
            animate={{ opacity: 1 }}
            transition={{ delay: 0.8 }}
            style={{ marginTop: '2rem', textAlign: 'center' }}
          >
            <p style={{ color: 'rgb(156, 163, 175)' }}>
              アカウントをお持ちでない方は{' '}
              <Link 
                to="/signup" 
                style={{ 
                  color: '#ffffff', 
                  fontWeight: '600', 
                  textDecoration: 'none',
                  transition: 'opacity 0.2s'
                }}
                onMouseEnter={(e) => e.currentTarget.style.opacity = '0.8'}
                onMouseLeave={(e) => e.currentTarget.style.opacity = '1'}
              >
                新規登録
              </Link>
            </p>
          </motion.div>
        </div>

        <motion.div
          initial={{ opacity: 0 }}
          animate={{ opacity: 1 }}
          transition={{ delay: 1 }}
          style={{ marginTop: '2rem', textAlign: 'center', fontSize: '0.875rem', color: 'rgb(107, 114, 128)' }}
        >
          <p>© 2025 omeroid Inc. All rights reserved.</p>
        </motion.div>
      </motion.div>
    </div>
  )
}