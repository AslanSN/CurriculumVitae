function copyText(id) {
	const text = document.getElementById(id).innerText;
	navigator.clipboard.writeText(text).then(() => {
		showCustomAlert(`${id === "mobile" ? "Mobile" : "Email"} added to clipboard!`)
	}).catch(err => {
		console.error(`Failed to copy ${id}: `, err);
	});
}

function closeCustomAlert() {
	customAlert = document.getElementById("custom-alert")
	if (customAlert !== null) {
		customAlert.style.display = 'none';
		customAlert.remove();
	}
}

function showCustomAlert(title, message = "") {
	const alertElement = document.createElement('div');
	alertElement.id = "custom-alert"
	alertElement.classList.add('alert', 'alert-success')
	alertElement.innerHTML = `
<svg width="48" height="48" viewBox="0 0 48 48" fill="none" xmlns="http://www.w3.org/2000/svg">
		<path fill-rule="evenodd" clip-rule="evenodd" d="M24 4C12.96 4 4 12.96 4 24C4 35.04 12.96 44 24 44C35.04 44 44 35.04 44 24C44 12.96 35.04 4 24 4ZM18.58 32.58L11.4 25.4C10.62 24.62 10.62 23.36 11.4 22.58C12.18 21.8 13.44 21.8 14.22 22.58L20 28.34L33.76 14.58C34.54 13.8 35.8 13.8 36.58 14.58C37.36 15.36 37.36 16.62 36.58 17.4L21.4 32.58C20.64 33.36 19.36 33.36 18.58 32.58Z" fill="#00BA34" />
	</svg>
	<div class="flex w-full justify-between items-center">
	<div class="flex flex-col">
		<span>${title}</span>
		<span class="text-content2">${message}</span>
	</div>
	<button id="close-alert-btn" class="btn btn-sm btn-circle btn-ghost">
<svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
			<path fill-rule="evenodd" clip-rule="evenodd" d="M18.3007 5.71C17.9107 5.32 17.2807 5.32 16.8907 5.71L12.0007 10.59L7.1107 5.7C6.7207 5.31 6.0907 5.31 5.7007 5.7C5.3107 6.09 5.3107 6.72 5.7007 7.11L10.5907 12L5.7007 16.89C5.3107 17.28 5.3107 17.91 5.7007 18.3C6.0907 18.69 6.7207 18.69 7.1107 18.3L12.0007 13.41L16.8907 18.3C17.2807 18.69 17.9107 18.69 18.3007 18.3C18.6907 17.91 18.6907 17.28 18.3007 16.89L13.4107 12L18.3007 7.11C18.6807 6.73 18.6807 6.09 18.3007 5.71Z" fill="#969696" />
		</svg>
	</button>
		</div>
`

	const alertContainer = document.getElementById("alert-container")
	alertContainer.appendChild(alertElement)

	document.getElementById('close-alert-btn').onclick = closeCustomAlert

	setTimeout(() => {
		closeCustomAlert();
	}, 4000)
}
