<template>
    <BaseSlide />
    
    <section id="about" class="py-16 px-4 sm:px-6 lg:px-8 bg-green-50">
    <h2 class="text-3xl font-semibold text-center text-indigo-600 mb-8">關於我們</h2>
    <p class="max-w-2xl mx-auto text-center">
      這裡是品牌介紹的詳細內容。描述您的品牌歷史、理念、價值觀等。讓潛在加盟者了解您的品牌故事和獨特之處。
    </p>
  </section>

  <section id="products" class="py-16 px-4 sm:px-6 lg:px-8">
    <div class="max-w-7xl mx-auto">
      <h2 class="text-3xl font-semibold text-center text-indigo-600 mb-8">產品與服務</h2>
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
        <div v-for="product in products" :key="product.title" class="bg-white rounded-lg shadow-md p-6 transition duration-300 hover:shadow-lg hover:-translate-y-1">
          <h3 class="text-xl font-semibold text-teal-600 mb-4">{{ product.title }}</h3>
          <p>{{ product.description }}</p>
        </div>
      </div>
    </div>
  </section>

  <section id="advantages" class="py-16 bg-gray-100 px-4 sm:px-6 lg:px-8">
    <div class="max-w-7xl mx-auto">
      <h2 class="text-3xl font-semibold text-center text-indigo-600 mb-8">加盟優勢</h2>
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
        <div v-for="advantage in advantages" :key="advantage.title" class="bg-white rounded-lg shadow-md p-6 transition duration-300 hover:shadow-lg hover:-translate-y-1">
          <h3 class="text-xl font-semibold text-teal-600 mb-4">{{ advantage.title }}</h3>
          <p>{{ advantage.description }}</p>
        </div>
      </div>
    </div>
  </section>

	<section id="process" class="py-16">
			<h2 class="text-3xl font-semibold text-center text-indigo-600 mb-8">加盟流程</h2>
			<div class="max-w-3xl mx-auto">
			<div v-for="(step, index) in steps" :key="index" class="flex mb-8">
					<div class="flex-shrink-0 w-10 h-10 rounded-full bg-indigo-600 flex items-center justify-center text-white font-bold">
					{{ index + 1 }}
					</div>
					<div class="ml-4">
					<h3 class="text-xl font-semibold mb-2">{{ step.title }}</h3>
					<p>{{ step.description }}</p>
					</div>
			</div>
			</div>
	</section>

	<section id="promotion" class="py-16 bg-indigo-100">
	<div class="max-w-4xl mx-auto px-4">
			<h2 class="text-3xl font-semibold text-center text-indigo-600 mb-8">線上展覽加盟優惠</h2>
			<div class="bg-white rounded-lg shadow-md p-8">
					<p class="text-xl mb-4">限時特惠！參與我們的線上展覽，享受特別的加盟優惠方案。包括：</p>
					<ul class="list-disc pl-5 mb-6 space-y-2">
					<li v-for="benefit in benefits" :key="benefit" class="text-lg">
							{{ benefit }}
					</li>
					</ul>
					<p class="text-lg font-semibold">優惠有效期至：{{ expirationDate }}</p>
			</div>
			</div>
	</section>

	<section id="contact" class="py-16">
		<div class="max-w-2xl mx-auto px-4">
			<h2 class="text-3xl font-semibold text-center text-indigo-600 mb-8">與我們聯繫</h2>
			<form @submit.prevent="submitForm" class="bg-white rounded-lg shadow-md p-8">
				<div class="mb-4">
					<label for="name" class="block text-gray-700 mb-2">姓名</label>
					<input type="text" id="name" v-model="form.name" required
									class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-indigo-500">
				</div>
				<div class="mb-4">
					<label for="email" class="block text-gray-700 mb-2">電子郵件</label>
					<input type="email" id="email" v-model="form.email" required
									class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-indigo-500">
				</div>
				<div class="mb-4">
					<label for="phone" class="block text-gray-700 mb-2">電話</label>
					<input type="tel" id="phone" v-model="form.phone"
									class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-indigo-500">
				</div>
				<div class="mb-6">
					<label for="message" class="block text-gray-700 mb-2">訊息</label>
					<textarea id="message" v-model="form.message" rows="4" required
										class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-indigo-500"></textarea>
				</div>
				<button type="submit"
								class="w-full bg-indigo-600 text-white py-2 px-4 rounded-md hover:bg-indigo-700 transition duration-300">
					發送訊息
				</button>
			</form>
		</div>
	</section>
</template>
  
<script>
import BaseSlide from './ui/BaseSlide.vue'

export default {
	name: 'BrandIntroduce',
	components: {
			BaseSlide
	},
	data() {
		return {
			products: [
					{ title: '產品/服務 1', description: '描述您的核心產品或服務，突出其特點和優勢。' },
					{ title: '產品/服務 2', description: '介紹另一個重要的產品或服務線，說明其市場需求和潛力。' },
					{ title: '產品/服務 3', description: '展示您的產品或服務組合的多樣性，吸引不同的目標客戶群。' }
			],
			advantages: [
					{ title: '品牌知名度', description: '詳細說明您的品牌在市場上的認知度和影響力。' },
					{ title: '營運支持', description: '介紹您為加盟商提供的全面支持，包括培訓、行銷等。' },
					{ title: '盈利模式', description: '解釋您的業務模式如何為加盟商創造持續的收益。' }
			],
			steps: [
					{ title: '提交加盟申請', description: '填寫並提交加盟申請表格。' },
					{ title: '初步溝通與評估', description: '我們的團隊將與您進行初步溝通，評估雙方的合作可能性。' },
					{ title: '實地考察與培訓', description: '安排實地考察，並提供初步的培訓。' },
					{ title: '簽署加盟合約', description: '雙方同意後，簽署正式的加盟合約。' },
					{ title: '開業準備與支持', description: '我們將全程協助您進行開業準備，並提供持續的支持。' }
			],
			benefits: [
					'首年加盟費減免50%',
					'免費獲得價值10萬元的開業支持包',
					'專屬的線上培訓課程',
					'額外的行銷推廣支持'
			],
			expirationDate: '2024年12月31日',
			form: {
				name: '',
				email: '',
				phone: '',
				message: ''
			}
		}
	},
	methods: {
		submitForm() {
			// 這裡處理表單提交邏輯
			console.log('表單提交', this.form)
			// 在實際應用中，您可能會發送 API 請求或執行其他操作
		}
	}
}
</script>
  