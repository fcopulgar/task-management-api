# 0008 — Usar bcrypt para hash de contraseñas

**Estado:** Aceptado  
**Fecha:** 2026-06-07

## Contexto

Las contraseñas deben almacenarse de forma segura y nunca en texto plano.

## Decisión

Usar bcrypt para hashear y verificar contraseñas.

## Consecuencias

bcrypt es una opcion probada para password hashing. Requiere gestionar costo adecuado y nunca exponer `password_hash`.

## Alternativas consideradas

Hash rápido general y Argon2. Hash rápido se descarta por inseguro para contraseñas; Argon2 queda como alternativa futura si se justifica.

## Relación con otros ADRs

Se relaciona con autenticación, cambio de contraseña y gestión de usuarios.
