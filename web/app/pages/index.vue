<template>
    <div>
        You are currently {{ status }}
    </div>

    <button @click="signout">Signout</button>

    <div v-if="data">
        <pre>{{ data }}</pre>
    </div>
    <div v-else>
        You are not logged in.
    </div>
</template>

<script setup lang="ts">
const { status, data, signOut, refresh } = useAuth()
const config = useRuntimeConfig()

const signout = async () => {
    const idToken = data.value?.idToken

    // 2. Clear local session first
    await signOut({ redirect: false })

    // 3. Redirect to Keycloak's logout endpoint
    const keycloakLogoutUrl = `${config.public.keycloakUrl}/realms/erp/protocol/openid-connect/logout`

    // Construct the URL with required params
    const url = new URL(keycloakLogoutUrl)
    url.searchParams.append('client_id', config.public.clientId)
    url.searchParams.append('post_logout_redirect_uri', window.location.origin) // Where to go after Keycloak logs out
    if (idToken) url.searchParams.append('id_token_hint', idToken)

    window.location.href = url.toString()

}

</script>