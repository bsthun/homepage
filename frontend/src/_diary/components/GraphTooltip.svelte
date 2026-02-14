<script lang="ts">
	import Tooltip, { Wrapper, Title, Content, Link, RichActions } from '@smui/tooltip'
	import { getContext, onMount } from 'svelte'
	import type { Writable } from 'svelte/store'

	const local: Writable<any> = getContext('local')

	let element: HTMLElement | null = null

	$: {
		if (element !== null && $local.activeGraph) {
			element.style.top = $local.activeGraph.event.clientY + 'px'
			element.style.left = $local.activeGraph.event.clientX + 'px'
			setTimeout(() => {
				element?.click()
			}, 0)
		}
	}
</script>

<div class="graph-tooltip">
	<Wrapper rich>
		<span class="target" bind:this={element} />
		<Tooltip persistent style="width: 120px">
			<Content>
				{$local.activeGraph?.day.title}
			</Content>
		</Tooltip>
	</Wrapper>
</div>

<style lang="scss">
	.graph-tooltip {
		position: absolute;
	}

	.target {
		position: fixed;
	}
</style>
