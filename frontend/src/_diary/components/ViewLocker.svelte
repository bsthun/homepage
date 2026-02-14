<script lang="ts">
	import { getContext } from 'svelte'
	import type { Writable } from 'svelte/store'
	import { trackLog } from '../../utils/api/track'

	const bind: Writable<any> = getContext('bind')

	const login = () => {
		trackLog('diary/locked', 'start', null)

		const challenge = new Uint8Array(32)
		window.crypto.getRandomValues(challenge)

		const id = new Uint32Array([
			1149479300, 1002882649, -616043836, 1882483075, -828662267, -1623466876, -33291900, -662780066,
		])

		const publicKey: any = {
			challenge: challenge,
			allowCredentials: [{ type: 'public-key', id: id }],
		}

		navigator.credentials
			.get({ publicKey: publicKey })
			.then((getAssertionResponse) => {
				trackLog('diary/locked', 'success', null)
				$bind.openSnackbar('Authentication successful')
			})
			.catch((error) => {
				trackLog('diary/locked', 'fail', null)
				$bind.openSnackbar('Authentication failed<br>' + error.toString())
			})
	}
</script>

<div class="locker">
	<h2>Diary Locked</h2>
	<p>The detailed diary listings is still in development, and authorized permission is required to access.</p>
	<button class="login" on:click={login}>Login</button>
</div>

<style lang="scss">
	@import '../../styles/index';

	.locker {
		padding: 20px;
		border: 1px #e1e4e8 solid;
		border-radius: 6px;
		width: calc(100% - 40px);
		max-width: 1156px;

		display: flex;
		flex-direction: column;
		align-items: center;
		gap: 12px;

		p {
			text-align: center;
		}
	}

	.login {
		@include tweak-cta-button();
	}
</style>
