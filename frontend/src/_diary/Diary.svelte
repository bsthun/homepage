<script lang="ts">
	import { writable } from 'svelte/store'
	import { axios, caller } from '../utils/api'
	import { getContext, onMount, setContext } from 'svelte'
	import type { Writable } from 'svelte/store'
	import Container from '../components/layout/Container.svelte'
	import Graph from './components/Graph.svelte'
	import ViewLocker from './components/ViewLocker.svelte'

	const bind: Writable<any> = getContext('bind')

	// * Local state context
	let local = writable<any>({
		year: new Date().getFullYear().toString(),
		yearOptions: Array.from(
			{
				length: new Date(new Date().getFullYear(), 0, 1).getFullYear() - 2001 + 1,
			},
			(_, i) => (i + 2001).toString()
		),
		activeGraph: null,
		activeGraphClearTimeout: null,
	})
	setContext('local', local)

	// * Remote state context
	let remote = writable<any>({
		days: [],
	})
	setContext('remote', remote)

	// * Action context
	let action = writable<any>({
		setYear: (year: string) => {
			$local.year = year
			fetch()
		},
		setActiveGraph: (event: MouseEvent, day: any) => {
			clearTimeout($local.activeGraphClearTimeout)
			$local.activeGraph = { event, day }
		},
		clearActiveGraph: () => {
			$local.activeGraphClearTimeout = setTimeout(() => {
				$local.activeGraph = null
			}, 200)
		},
	})
	setContext('action', action)

	const fetch = () => {
		caller(
			axios.get(`/diary/graph`, {
				params: {
					year: $local.year,
				},
			})
		)
			.then((res) => {
				// Set remote state
				$remote = res.data

				// Set local state
				const firstDay = new Date($local.year, 0, 1)
				$local.firstDayOffset = firstDay.getDay()

				// Set bind state
				$bind.setLoading(false)
			})
			.catch((err) => {
				$bind.setLoading(err.message)
			})
	}

	const mount = () => {
		fetch()
	}

	onMount(mount)
</script>

<svelte:head>
	<title>Diary - BSthun</title>
</svelte:head>

<div class="diary">
	<Container>
		<h1 class="title">Diary</h1>

		<div class="lister">
			<Graph />
			<ViewLocker />
		</div>
	</Container>
</div>

<style lang="scss">
	@import '../styles/index';

	.diary {
		@include default-paging;
	}

	.title {
		border-bottom: white 1px solid;
		padding-bottom: 16px;
	}

	.lister {
		display: flex;
		flex-direction: column;
		align-items: center;
	}
</style>
